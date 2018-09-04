import {api, IPageData} from "./api";
import {read, write} from "./storage";
import {show} from "./error";

type callback<T> = (result: T) => void;

export interface IClient {
    page: {
        create(email: string, fn: callback<string>);
        delete(addr: string, fn: callback<undefined>);
        edit(addr: string, spec: string, fn: callback<IPageData>);
        fetch(addr: string, fn: callback<IPageData>);
        publish(addr: string, fn: callback<undefined>);
        validate(spec: string, fn: callback<undefined>);
        password: {
            change(addr: string, pass: string, fn: callback<undefined>);
            reset(addr: string, email: string, fn: callback<undefined>);
        };
        token: {
            create(addr: string, pass: string, fn: callback<string>);
        };
    };
}

class TokenError extends Error {
    constructor(addr: string) {
        super(`Missing access token for address ${addr}`);
    }
}

const catchAll = <T>(value: T): T => {
    if (typeof value === "function") {
        const replacement = async (...a: any[]) => {
            try {
                return await value(...a);
            } catch (e) {
                show(e.toString());
            }
        };
        return replacement as any as T;
    }
    if (typeof value !== "object") {
        return value;
    }
    const copy = {};
    Object.keys(value).forEach((key) => {
        copy[key] = catchAll(value[key]);
    });
    return copy as T;
};

export const client: IClient = catchAll({
    page: {
        create: async (email, fn) => {
            const res = await api.page.create(email);
            fn(res);
        },
        delete: async (addr, fn) => {
            const {token} = read(addr);
            if (token === null) {
                throw new TokenError(addr);
            }
            await api.page.delete(addr, token);
            fn(undefined);
        },
        edit: async (addr, spec, fn) => {
            const {token} = read(addr);
            if (token === null) {
                throw new TokenError(addr);
            }
            const data = await api.page.edit(addr, token, spec);
            write(addr, {data});
            fn(data);
        },
        fetch: async (addr, fn) => {
            const {data: staleData, token} = read(addr);
            if (staleData !== null) {
                fn(staleData);
            }
            const data = await api.page.fetch(addr, token || undefined);
            write(addr, {data});
            fn(data);

        },
        publish: async (addr, fn) => {
            const {token} = read(addr);
            if (token === null) {
                throw new TokenError(addr);
            }
            await api.page.publish(addr, token);
            fn(undefined);
        },
        validate: async (spec, fn) => {
            await api.page.validate(spec);
            fn(undefined);
        },
        password: {
            change: async (addr, pass, fn) => {
                const {token} = read(addr);
                if (token === null) {
                    throw new TokenError(addr);
                }
                await api.page.password.change(addr, token, pass);
                fn(undefined);
            },
            reset: async (addr, email, fn) => {
                await api.page.password.reset(addr, email);
                fn(undefined);
            },
        },
        token: {
            create: async (addr, pass, fn) => {
                const token = await api.page.token.create(addr, pass);
                write(addr, {token});
                fn(token);
            },
        },
    },
} as IClient);

import {api, IPageData} from "./api";
import {read, write} from "./storage";
import {show} from "./error";

export {IPageData} from "./api";

type operation<T = any, S extends any[] = any[]> = (
    callback: (result: T) => void,
    errorHandler: (message: string) => void | null,
    ...args: S
) => void;

export interface IClient {
    page: {
        create: operation<string, [string]>;
        delete: operation<undefined, [string]>;
        edit: operation<IPageData, [string, string]>;
        fetch: operation<IPageData, [string]>;
        publish: operation<undefined, [string]>;
        validate: operation<undefined, [string]>;
        password: {
            change: operation<undefined, [string, string]>;
            reset: operation<undefined, [string, string]>;
        };
        token: {
            create: operation<string, [string, string]>;
        };
    };
}

const missingTokenMessage = (name: string, addr: string): string => {
    return `Missing access token for operation "${name}" on address "${addr}"`;
};

// Wrap all the client's operations to handle errors.
const catchAll = <T>(value: T): T => {
    if (typeof value === "function") {
        const op = value as any as operation;
        const replacement: operation = async (cb, handler = show, ...args) => {
            try {
                await op(cb, handler, ...args);
            } catch (e) {
                handler(e.toString());
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
        create: async (callback, _, email) => {
            const res = await api.page.create(email);
            callback(res);
        },
        delete: async (callback, errorHandler, addr) => {
            const {token} = read(addr);
            if (token === null) {
                return errorHandler(missingTokenMessage("delete", addr));
            }
            await api.page.delete(addr, token);
            callback(undefined);
        },
        edit: async (callback, errorHandler, addr, spec) => {
            const {token} = read(addr);
            if (token === null) {
                return errorHandler(missingTokenMessage("edit", addr));
            }
            const data = await api.page.edit(addr, token, spec);
            write(addr, {data});
            callback(data);
        },
        fetch: async (callback, _, addr) => {
            const {data: staleData, token} = read(addr);
            if (staleData !== null) {
                callback(staleData);
            }
            const data = await api.page.fetch(addr, token || undefined);
            write(addr, {data});
            callback(data);

        },
        publish: async (callback, errorHandler, addr) => {
            const {token} = read(addr);
            if (token === null) {
                return errorHandler(missingTokenMessage("publish", addr));
            }
            await api.page.publish(addr, token);
            callback(undefined);
        },
        validate: async (callback, _, spec) => {
            await api.page.validate(spec);
            callback(undefined);
        },
        password: {
            change: async (callback, errorHandler, addr, pass) => {
                const {token} = read(addr);
                if (token === null) {
                    return errorHandler(missingTokenMessage("change password", addr));
                }
                await api.page.password.change(addr, token, pass);
                callback(undefined);
            },
            reset: async (callback, _, addr, email) => {
                await api.page.password.reset(addr, email);
                callback(undefined);
            },
        },
        token: {
            create: async (callback, _, addr, pass) => {
                const token = await api.page.token.create(addr, pass);
                write(addr, {token});
                callback(token);
            },
        },
    },
} as IClient);

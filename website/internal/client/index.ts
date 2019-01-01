import {api} from "./api";
import {read, write} from "./storage";
import {show} from "./error";

export {IPageData} from "./types";

type clientFunc<T = any, S extends any[] = any[]> = (
    callback: (result: T) => void,
    errorHandler: ((message: string) => void) | null,
    ...args: S
) => void;

// prettier-ignore
export type Client<T> = {
    [K in keyof T]:
        T[K] extends (...args: infer A) => infer R
            ? (R extends Promise<infer D>
                ? clientFunc<D, A>
                : never)
            : Client<T[K]>;
};

const missingTokenMessage = (name: string, addr: string): string => {
    return `Missing access token for operation "${name}" on address "${addr}"`;
};

// Wrap all the client's operations to handle errors.
const catchAll = <T>(value: T): T => {
    if (typeof value === "function") {
        const op = (value as any) as clientFunc;
        const replacement: clientFunc = async (cb, handler = show, ...args) => {
            try {
                await op(cb, handler, ...args);
            } catch (e) {
                handler(e.toString());
            }
        };
        return (replacement as any) as T;
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

// const a: Client<typeof api> = {
//     page: {
//         create: (cb, eh, a) => {},
//     },
// };

export const client: Client<typeof api> = catchAll({
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
                    return errorHandler(
                        missingTokenMessage("change password", addr),
                    );
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
});

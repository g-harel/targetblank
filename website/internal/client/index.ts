import {api} from "./api";
import {read, write} from "./storage";
import {show} from "./error";
import {IPageData} from "./types";

export {IPageData} from "./types";

type Callback<T> = (value: T) => void;
type ErrorHandler = Callback<string> | null;

const missingTokenMessage = (name: string, addr: string): string => {
    return `Missing access token for operation "${name}" on address "${addr}"`;
};

export const client = {
    page: {
        create: (cb: Callback<string>, email: string) => {
            api.page.create(email).then(cb);
        },
        delete: (cb: Callback<void>, err: ErrorHandler, addr: string) => {
            const {token} = read(addr);
            if (token === null) {
                return err(missingTokenMessage("delete", addr));
            }
            api.page.delete(addr, token).then(cb);
        },
        edit: (
            cb: Callback<IPageData>,
            err: ErrorHandler,
            addr: string,
            spec: string,
        ) => {
            const {token} = read(addr);
            if (token === null) {
                return err(missingTokenMessage("edit", addr));
            }
            api.page
                .edit(addr, token, spec)
                .then((data) => {
                    write(addr, {data});
                    return data;
                })
                .then(cb);
        },
        fetch: (cb: Callback<IPageData>, addr: string) => {
            const {data: staleData, token} = read(addr);
            if (staleData !== null) {
                cb(staleData);
            }
            api.page
                .fetch(addr, token || undefined)
                .then((data) => {
                    write(addr, {data});
                    return data;
                })
                .then(cb);
        },
        publish: (cb: Callback<void>, err: ErrorHandler, addr: string) => {
            const {token} = read(addr);
            if (token === null) {
                return err(missingTokenMessage("publish", addr));
            }
            api.page.publish(addr, token).then(cb);
        },
        validate: (cb: Callback<void>, spec: string) => {
            api.page.validate(spec).then(cb);
        },
        password: {
            change: (cb: Callback<void>, err: ErrorHandler, addr, pass) => {
                const {token} = read(addr);
                if (token === null) {
                    return err(missingTokenMessage("change password", addr));
                }
                api.page.password.change(addr, token, pass).then(cb);
            },
            reset: (cb: Callback<void>, addr, email) => {
                api.page.password.reset(addr, email).then(cb);
            },
        },
        token: {
            create: (cb: Callback<string>, addr: string, pass: string) => {
                api.page.token
                    .create(addr, pass)
                    .then((token) => {
                        write(addr, {token});
                        return token;
                    })
                    .then(cb);
            },
        },
    },
};

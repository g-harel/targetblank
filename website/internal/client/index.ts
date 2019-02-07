import * as api from "../api";
import {IPageData} from "../types";
import {remoteClient} from "./remote";

export {IPageData} from "../types";

export const localAddr = "local";

export type Callback<T = void> = (value: T) => void;
export type ErrHandler = Callback<string>;

export interface Client {
    isAuthorized(): boolean;
    pageDelete(cb: Callback, err: ErrHandler): void;
    pageUpdate(cb: Callback<IPageData>, err: ErrHandler, doc: string): void;
    pageRead(cb: Callback<IPageData>, err: ErrHandler): void;
    passUpdate(
        cb: Callback,
        err: ErrHandler,
        pass: string,
        token?: string,
    ): void;
    passReset(cb: Callback, err: ErrHandler, email: string): void;
    tokenCreate(cb: Callback<string>, err: ErrHandler, pass: string): void;
}

interface ClientGenerator {
    (addr: string): Client;
}

interface StaticClient {
    pageCreate: (cb: Callback<string>, err: ErrHandler, email: string) => void;
    pageValidate: (cb: Callback<void>, err: ErrHandler, doc: string) => void;
}

const clientGenerator: ClientGenerator = (addr) => {
    if (addr !== localAddr) {
        return remoteClient(addr);
    }
    throw "can't use local address";
};

const staticClient: StaticClient = {
    pageCreate: (cb, err, email) => {
        api.pageCreate(email)
            .then(cb)
            .catch(err);
    },
    pageValidate: (cb, err, doc) => {
        api.pageValidate(doc)
            .then(cb)
            .catch(err);
    },
};

export const client = Object.assign(clientGenerator, staticClient);

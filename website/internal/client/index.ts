import * as api from "../api";
import {IPageData} from "../types";
import {RemoteClient} from "./remote";

export {IPageData} from "../types";

export const localAddr = "local";

type Callback<T = void> = (value: T) => void;
type ErrHandler = Callback<string> | null;

export abstract class Client {
    protected addr: string;

    constructor(addr: string) {
        this.addr = addr;
    }

    abstract isAuthorized(): boolean;
    abstract pageDelete(cb: Callback, err: ErrHandler): void;
    abstract pageUpdate(cb: Callback<IPageData>, err: ErrHandler, doc: string);
    abstract pageRead(cb: Callback<IPageData>, err: ErrHandler);
    abstract passUpdate(
        cb: Callback,
        err: ErrHandler,
        pass: string,
        token?: string,
    );
    abstract passReset(cb: Callback, err: ErrHandler, email: string);
    abstract tokenCreate(cb: Callback<string>, err: ErrHandler, pass: string);
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
        return new RemoteClient(addr);
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

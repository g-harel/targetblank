import * as api from "../api";
import {read, write} from "../storage";
import {IPageData} from "../types";

export {IPageData} from "../types";

type Callback<T> = (value: T) => void;
type ErrorHandler = Callback<string> | null;

const missingToken = (name: string, addr: string): string => {
    console.error(`Missing access token for "${name}" on address "${addr}"`);
    return "Unauthorized";
};

const writeToken = (addr: string) => (token: string) => {
    write(addr, {token});
    return token;
};

const writeData = (addr: string) => (data: IPageData) => {
    write(addr, {data});
    return data;
};

class PageClient {
    public password: PagePasswordClient;
    public token: PageTokenClient;

    constructor() {
        this.password = new PagePasswordClient();
        this.token = new PageTokenClient();
    }

    auth(addr: string): boolean {
        return !!read(addr).token;
    }

    create(cb: Callback<string>, err: ErrorHandler, email: string) {
        api.pageCreate(email)
            .then(cb)
            .catch(err);
    }

    delete(cb: Callback<void>, err: ErrorHandler, addr: string) {
        const {token} = read(addr);
        if (token === null) {
            return err(missingToken("delete", addr));
        }
        api.pageDelete(addr, token)
            .then(cb)
            .catch(err);
    }

    edit(
        cb: Callback<IPageData>,
        err: ErrorHandler,
        addr: string,
        doc: string,
    ) {
        const {token} = read(addr);
        if (token === null) {
            return err(missingToken("edit", addr));
        }
        api.pageUpdate(addr, token, doc)
            .then(writeData(addr))
            .then(cb)
            .catch(err);
    }

    fetch(cb: Callback<IPageData>, err: ErrorHandler, addr: string) {
        const {data: cachedData, token} = read(addr);
        if (cachedData !== null) cb(cachedData);
        api.pageRead(addr, token || undefined)
            .then(writeData(addr))
            .then(cb)
            .catch(err);
    }

    validate(cb: Callback<void>, err: ErrorHandler, doc: string) {
        api.pageValidate(doc)
            .then(cb)
            .catch(err);
    }
}

class PagePasswordClient {
    change(
        cb: Callback<void>,
        err: ErrorHandler,
        addr: string,
        pass: string,
        token?: string,
    ) {
        const t = token || read(addr).token;
        if (t === null) {
            return err(missingToken("change password", addr));
        }
        api.passwordUpdate(addr, t, pass)
            .then(cb)
            .catch(err);
    }

    reset(cb: Callback<void>, err: ErrorHandler, addr: string, email: string) {
        api.passwordReset(addr, email)
            .then(cb)
            .catch(err);
    }
}

class PageTokenClient {
    create(
        cb: Callback<string>,
        err: ErrorHandler,
        addr: string,
        pass: string,
    ) {
        api.tokenCreate(addr, pass)
            .then(writeToken(addr))
            .then(cb)
            .catch(err);
    }
}

export const client = {
    page: new PageClient(),
};

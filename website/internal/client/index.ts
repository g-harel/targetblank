import {PageTokenAPI, PagePasswordAPI, PageAPI} from "./api";
import {read, write} from "./storage";
import {IPageData} from "./types";

export {IPageData} from "./types";

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
    private api: PageAPI;

    public password: PagePasswordClient;
    public token: PageTokenClient;

    constructor(api: PageAPI) {
        this.api = api;
        this.password = new PagePasswordClient(new PagePasswordAPI());
        this.token = new PageTokenClient(new PageTokenAPI());
    }

    create(cb: Callback<string>, err: ErrorHandler, email: string) {
        this.api
            .create(email)
            .then(cb)
            .catch(err);
    }

    delete(cb: Callback<void>, err: ErrorHandler, addr: string) {
        const {token} = read(addr);
        if (token === null) {
            return err(missingToken("delete", addr));
        }
        this.api
            .delete(addr, token)
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
        this.api
            .edit(addr, token, doc)
            .then(writeData(addr))
            .then(cb)
            .catch(err);
    }

    fetch(cb: Callback<IPageData>, err: ErrorHandler, addr: string) {
        const {data: cachedData, token} = read(addr);
        if (cachedData !== null) cb(cachedData);
        this.api
            .fetch(addr, token || undefined)
            .then(writeData(addr))
            .then(cb)
            .catch(err);
    }

    publish(cb: Callback<void>, err: ErrorHandler, addr: string) {
        const {token} = read(addr);
        if (token === null) {
            return err(missingToken("publish", addr));
        }
        this.api
            .publish(addr, token)
            .then(cb)
            .catch(err);
    }

    validate(cb: Callback<void>, err: ErrorHandler, doc: string) {
        this.api
            .validate(doc)
            .then(cb)
            .catch(err);
    }
}

class PagePasswordClient {
    private api: PagePasswordAPI;

    constructor(api: PagePasswordAPI) {
        this.api = api;
    }

    change(cb: Callback<void>, err: ErrorHandler, addr, pass) {
        const {token} = read(addr);
        if (token === null) {
            return err(missingToken("change password", addr));
        }
        this.api
            .change(addr, token, pass)
            .then(cb)
            .catch(err);
    }

    reset(cb: Callback<void>, err: ErrorHandler, addr: string, email: string) {
        this.api
            .reset(addr, email)
            .then(cb)
            .catch(err);
    }
}

class PageTokenClient {
    private api: PageTokenAPI;

    constructor(api: PageTokenAPI) {
        this.api = api;
    }

    create(
        cb: Callback<string>,
        err: ErrorHandler,
        addr: string,
        pass: string,
    ) {
        this.api
            .create(addr, pass)
            .then(writeToken(addr))
            .then(cb)
            .catch(err);
    }
}

export const client = {
    page: new PageClient(new PageAPI()),
};

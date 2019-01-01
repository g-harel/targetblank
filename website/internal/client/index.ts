import {PageTokenAPI, PagePasswordAPI, PageAPI} from "./api";
import {read, write} from "./storage";
import {IPageData} from "./types";

export {IPageData} from "./types";

type Callback<T> = (value: T) => void;
type ErrorHandler = Callback<string> | null;

const missingTokenMessage = (name: string, addr: string): string => {
    return `Missing access token for operation "${name}" on address "${addr}"`;
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

    create(cb: Callback<string>, email: string) {
        this.api.create(email).then(cb);
    }

    delete(cb: Callback<void>, err: ErrorHandler, addr: string) {
        const {token} = read(addr);
        if (token === null) {
            return err(missingTokenMessage("delete", addr));
        }
        this.api.delete(addr, token).then(cb);
    }

    edit(
        cb: Callback<IPageData>,
        err: ErrorHandler,
        addr: string,
        spec: string,
    ) {
        const {token} = read(addr);
        if (token === null) {
            return err(missingTokenMessage("edit", addr));
        }
        this.api
            .edit(addr, token, spec)
            .then(writeData(addr))
            .then(cb);
    }

    fetch(cb: Callback<IPageData>, addr: string) {
        const {data: cachedData, token} = read(addr);
        if (cachedData !== null) cb(cachedData);
        this.api
            .fetch(addr, token || undefined)
            .then(writeData(addr))
            .then(cb);
    }

    publish(cb: Callback<void>, err: ErrorHandler, addr: string) {
        const {token} = read(addr);
        if (token === null) {
            return err(missingTokenMessage("publish", addr));
        }
        this.api.publish(addr, token).then(cb);
    }

    validate(cb: Callback<void>, spec: string) {
        this.api.validate(spec).then(cb);
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
            return err(missingTokenMessage("change password", addr));
        }
        this.api.change(addr, token, pass).then(cb);
    }

    reset(cb: Callback<void>, addr: string, email: string) {
        this.api.reset(addr, email).then(cb);
    }
}

class PageTokenClient {
    private api: PageTokenAPI;

    constructor(api: PageTokenAPI) {
        this.api = api;
    }

    create(cb: Callback<string>, addr: string, pass: string) {
        this.api
            .create(addr, pass)
            .then(writeToken(addr))
            .then(cb);
    }
}

export const client = {
    page: new PageClient(new PageAPI()),
};

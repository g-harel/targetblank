import * as api from "../api";
import {read, write} from "../storage";
import {IPageData} from "../types";
import {Client} from ".";

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

export class RemoteClient extends Client {
    isAuthorized() {
        return !!read(this.addr).token;
    }

    pageDelete(cb, err) {
        const {token} = read(this.addr);
        if (token === null) {
            return err(missingToken("delete", this.addr));
        }
        api.pageDelete(this.addr, token)
            .then(cb)
            .catch(err);
    }

    pageUpdate(cb, err, doc) {
        const {token} = read(this.addr);
        if (token === null) {
            return err(missingToken("edit", this.addr));
        }
        api.pageUpdate(this.addr, token, doc)
            .then(writeData(this.addr))
            .then(cb)
            .catch(err);
    }

    pageRead(cb, err) {
        const {data: cachedData, token} = read(this.addr);
        if (cachedData !== null) cb(cachedData);
        api.pageRead(this.addr, token || undefined)
            .then(writeData(this.addr))
            .then(cb)
            .catch(err);
    }

    passUpdate(cb, err, pass, token) {
        const t = token || read(this.addr).token;
        if (t === null) {
            return err(missingToken("change password", this.addr));
        }
        api.passUpdate(this.addr, t, pass)
            .then(cb)
            .catch(err);
    }

    passReset(cb, err, email) {
        api.passReset(this.addr, email)
            .then(cb)
            .catch(err);
    }

    tokenCreate(cb, err, pass) {
        api.tokenCreate(this.addr, pass)
            .then(writeToken(this.addr))
            .then(cb)
            .catch(err);
    }
}

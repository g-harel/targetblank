import * as api from "../api";
import {read, write} from "../storage";
import {IPageData} from "../types";
import {Client} from ".";

const missingToken = (name: string, addr: string): api.IRequestError => {
    console.error(`Missing access token for "${name}" on address "${addr}"`);
    return {message: "Unauthorized", isAuth: true};
};

const writeToken = (addr: string) => (token: string) => {
    write(addr, {token});
    return token;
};

const writeData = (addr: string) => (data: IPageData) => {
    write(addr, {data});
    return data;
};

export const remoteClient = (addr: string): Client => ({
    isAuthorized() {
        return !!read(addr).token;
    },

    pageUpdate(cb, err, doc) {
        const {token} = read(addr);
        if (token === null) {
            return err(missingToken("edit", addr));
        }
        api.pageUpdate(addr, token, doc)
            .then(writeData(addr))
            .then(cb)
            .catch(err);
    },

    pageRead(cb, err) {
        const {data: cachedData, token} = read(addr);
        if (cachedData !== null) cb(cachedData);
        api.pageRead(addr, token || "")
            .then(writeData(addr))
            .then(cb)
            .catch(err);
    },

    passUpdate(cb, err, pass, token) {
        const t = token || read(addr).token;
        if (t === null) {
            return err(missingToken("change password", addr));
        }
        api.passUpdate(addr, t, pass)
            .then(cb)
            .catch(err);
    },

    passReset(cb, err, email) {
        api.passReset(addr, email)
            .then(cb)
            .catch(err);
    },

    tokenCreate(cb, err, pass) {
        api.tokenCreate(addr, pass)
            .then(writeToken(addr))
            .then(cb)
            .catch(err);
    },
});

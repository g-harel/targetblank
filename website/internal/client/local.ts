import * as api from "../api";
import {read, write} from "../storage";
import {IPageData} from "../types";
import {Client, localAddr} from ".";

const defaultPage = `version 1\n===`;

const writeData = (data: IPageData) => {
    write(localAddr, {data});
    return data;
};

export const localClient = (): Client => ({
    isAuthorized() {
        return true;
    },

    pageUpdate(cb, err, doc) {
        api.pageValidate(doc)
            .then(writeData)
            .then(cb)
            .catch(err);
    },

    pageRead(cb, err) {
        const {data} = read(localAddr);
        if (data) {
            return cb(data);
        }

        api.pageValidate(defaultPage)
            .then(writeData)
            .then(cb)
            .catch(err);
    },

    passUpdate(cb) {
        cb();
    },

    passReset(cb) {
        cb();
    },

    tokenCreate(cb) {
        cb("");
    },
});

import * as api from "../api";
import {read, write} from "../storage";
import {IPageData} from "../types";
import {Client, localAddr} from ".";

const defaultPage = `version 1
title = Welcome!
===
This is your local page, it behaves exactly like a regular page, but it's for your eyes only!
The contents are only saved on this browser, so they may get reset when clearing browser data.
Click the edit button in the top right to start making changes.
---
Helpful Links
    Document Format [https://github.com/g-harel/targetblank/#document-format]
    Keyboard Shortcuts [https://github.com/g-harel/targetblank/#keyboard-shortcuts]
`;

const writeData = (data: IPageData) => {
    write(localAddr, {data});
    return data;
};

export const localClient = (): Client => ({
    resetAuth() {},

    isAuthorized() {
        return true;
    },

    pageUpdate(cb, err, doc) {
        api.pageValidate(doc).then(writeData).then(cb).catch(err);
    },

    pageRead(cb, err) {
        const {data} = read(localAddr);
        if (data) {
            return cb(data);
        }

        api.pageValidate(defaultPage).then(writeData).then(cb).catch(err);
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

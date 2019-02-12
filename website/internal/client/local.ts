import * as api from "../api";
import {read, write} from "../storage";
import {IPageData} from "../types";
import {Client, localAddr} from ".";

// tslint:disable-next-line:max-line-length
const defaultPage = `# Everything after a pound character (#), trailing whitespace and empty lines are ignored.

# Documents must start with their version (currently only 1).
version 1

# Document metadata key-value pairs can be added at the top of the document.
key=value

# The "title" key can be used to name the document.
title=Hello World

# The first group starts after the header line.
===

# Group metadata key-value pairs can be added at the start of each group.
# These values are currently ignored, but may be used in the future.
key=value

# Groups contain entries containing a label and a link.
labelled link [example.com]

# Both the label and the link are optional.
label without link
[google.com]
amazon.com

# New groups are started using the group delimiter.
---

# Group entries can be nested indefinitely (using indentation).
entry 1
    entry 2
        entry 3
    entry 4
`;

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

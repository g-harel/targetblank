import {IPageData} from "./types";

// Tokens expire after three days.
const tokenTTL = 1000 * 60 * 60 * 24 * 3;

type Cache = {
    token: string | null;
    data: IPageData | null;
};

type Expiry = {
    time: number;
};

// Generates local storage key for given address and data type.
const keyData = (addr: string) => `addr:${addr}`;
const keyExpiry = (addr: string) => `kill:${addr}`;

// Generates a zeroed-out page cache value.
const empty = () => ({token: null, data: null});

// Read page data from local storage.
export const read = (addr: string): Cache => {
    const data: Cache | null = JSON.parse(localStorage.getItem(keyData(addr)));
    if (!data) {
        return empty();
    }

    // Remove expired tokens.
    if (data.token) {
        const expiry: Expiry | null = JSON.parse(localStorage.getItem(keyExpiry(addr)));
        if (!expiry || expiry.time < Date.now()) {
            write(addr, {token: null});
            data.token = null;
        }
    }

    return data;
};

// Update stored data from local storage.
export const write = (addr: string, values: Partial<Cache>) => {
    const data: Cache = JSON.parse(localStorage.getItem(keyData(addr))) || empty();

    // Update expiry time when token is written.
    if (values.token && values.token !== data.token) {
        const e: Expiry = {time: Date.now() + tokenTTL};
        localStorage.setItem(keyExpiry(addr), JSON.stringify(e));
    }

    Object.assign(data || empty(), values);
    localStorage.setItem(keyData(addr), JSON.stringify(data));
};

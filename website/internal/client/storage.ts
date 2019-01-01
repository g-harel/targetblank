import {IPageData} from "./types";

type Cache = {
    token: string | null;
    data: IPageData | null;
};

// Generates local storage key for given address.
const key = (addr: string) => {
    return `addr:${addr}`;
};

// Generates a zeroed-out page cache value.
const empty = () => ({token: null, data: null});

// Read page data from local storage.
export const read = (addr: string): Cache => {
    const data: Cache | null = JSON.parse(localStorage.getItem(key(addr)));
    return data || empty();
};

// Update stored data from local storage.
export const write = (addr: string, values: Partial<Cache>) => {
    const data: Cache = JSON.parse(localStorage.getItem(key(addr))) || empty();
    Object.assign(data || empty(), values);
    localStorage.setItem(key(addr), JSON.stringify(data));
};

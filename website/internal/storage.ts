import {IPageData} from "./types";

type Cache = {
    token: string | null;
    data: IPageData | null;
};

// Generates local storage key for given address and data type.
const keyData = (addr: string) => `addr:${addr}`;

// Generates a zeroed-out page cache value.
const empty = (): Cache => ({token: null, data: null});

// Read page data from local storage.
export const read = (addr: string): Cache => {
    const data: Cache | null = JSON.parse(
        localStorage.getItem(keyData(addr)) || "null",
    );
    if (!data) {
        return empty();
    }

    return data;
};

// Update stored data from local storage.
export const write = (addr: string, values: Partial<Cache>) => {
    const data: Cache =
        JSON.parse(localStorage.getItem(keyData(addr)) || "null") || empty();

    Object.assign(data || empty(), values);
    localStorage.setItem(keyData(addr), JSON.stringify(data));
};

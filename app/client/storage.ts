import {IPageData} from "./api";

type stored = {
    token: string | null,
    data: IPageData | null,
};

// Generates local storage key for given address.
const key = (addr: string) => {
    return "addr:" + addr;
};

// Generates a zeroed-out stored value.
const empty = () => ({token: null, data: null});

// Read page data from local storage.
export const read = (addr: string): stored => {
    const data: stored | null = JSON.parse(localStorage.getItem(key(addr)));
    return data || empty();
};

// Update stored data from local storage.
export const write = (addr: string, values: Partial<stored>) => {
    const data: stored = JSON.parse(localStorage.getItem(key(addr))) || empty();
    Object.assign(data || empty(), values);
    localStorage.setItem(key(addr), JSON.stringify(data));
};

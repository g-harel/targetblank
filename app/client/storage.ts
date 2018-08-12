import {IPageData} from "./api";

const key = "targetblank-store";

type stored = {
    token: string | null,
    data: IPageData | null,
};

export interface IStore {
    [address: string]: stored;
}

const empty = () => ({token: null, data: null});

if (!localStorage.getItem(key)) {
    localStorage.setItem(key, "{}");
}

export const read = (addr: string): stored => {
    const data = JSON.parse(localStorage.getItem(key)) as IStore;
    return data[addr] || empty();
};

export const save = (addr: string, s: Partial<stored>) => {
    const data = JSON.parse(localStorage.getItem(key)) as IStore;
    data[addr] = Object.assign(data[addr] || empty(), s);
    localStorage.setItem(key, JSON.stringify(data));
};

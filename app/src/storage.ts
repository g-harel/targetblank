import {fetchPage} from "./api";

const key = "targetblank-store";

interface IStore {
    [address: string]: {
        token: string,
        data: IPageData,
    };
}

const readStore = (): IStore => {
    return JSON.parse(localStorage.getItem(key));
};

const saveStore = (s: IStore) => {
    localStorage.setItem(key, JSON.stringify(s));
    return s;
};

const updateStore = async (u: (s: IStore) => IStore) => {
    const s = await u(readStore());
    saveStore(s);
    return s;
};

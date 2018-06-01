const key = "targetblank-store";

export interface IStore {
    [address: string]: {
        token: string,
        data: IPageData,
    };
}

export const read = (): IStore => {
    return JSON.parse(localStorage.getItem(key));
};

export const save = (s: IStore) => {
    localStorage.setItem(key, JSON.stringify(s));
    return s;
};

export const update = async (u: (s: IStore) => IStore) => {
    const s = await u(read());
    save(s);
    return s;
};

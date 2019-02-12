export type Component<P extends object = {}, D = undefined> = (
    params: P,
    update: (data?: D) => void,
) => any;

export interface IPageEntry {
    label: string;
    link: string;
    entries: IPageEntry[];
}

export interface IPageGroup {
    meta: {
        [key: string]: string;
    };
    entries: IPageEntry[];
}

export interface IPageData {
    version: string;
    raw: string;
    meta: {
        [key: string]: string;
    };
    groups: IPageGroup[];
}

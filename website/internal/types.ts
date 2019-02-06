export type Component<P extends object = {}, D = undefined> = (
    params: P,
    update: (data?: D) => void,
) => any;

export interface IPageItem {
    label: string;
    link: string;
    entries: IPageItem[];
}

export interface IPageGroup {
    meta: {
        [key: string]: string;
    };
    entries: IPageItem[];
}

export interface IPageData {
    version: string;
    raw: string;
    meta: {
        [key: string]: string;
    };
    groups: IPageGroup[];
}

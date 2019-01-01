export interface IPageItem {
    label: string;
    link: string;
    items: IPageItem[];
}

export interface IPageGroup {
    meta: {
        [key: string]: string;
    };
    items: IPageItem[];
}

export interface IPageData {
    version: string;
    spec: string;
    meta: {
        [key: string]: string;
    };
    groups: IPageGroup[];
}

export interface IError {
    statusCode: number;
    message: string;
}

interface IPageItem {
    label: string;
    link: string;
    items: IPageItem[];
}

interface IPageGroup {
    meta: {
        [key: string]: string;
    };
    items: IPageItem[];
}

interface IPageData {
    version: string;
    spec: string;
    meta: {
        [key: string]: string;
    };
    groups: IPageGroup[];
}

declare module "okwolo/lite" {
    export interface App<S> {
        (...any): void

        setState(state: S): void
        setState(updater: (state: S) => S): void

        getState(): S

        use(blob: string, ...any): void
    }

    export default function<S>(target?: any, global?: any): App<S>
}
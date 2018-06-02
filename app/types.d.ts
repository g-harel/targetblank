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
        (f: () => (state: S) => Element): void
        <P extends object>(path: string | RegExp, f: (params: P) => (state: S) => Element): void

        setState(state: NonNullable<S>): void
        setState(updater: (state: S) => NonNullable<S>): void

        getState(): S

        use(blob: string, ...any): void
    }

    export type BlankElement = boolean | null;

    export type TextElement = string | number;

    export type TagElement = {
        0: string;
        1?: object;
        2?: Element[];
    } & any[]

    export type ComponentElement<P extends object = {}, A extends object = {}> = {
        0: (params: P) => (arg: A | undefined) => Element;
        1?: P;
        2?: Element[];
    } & any[]

    export type Element = BlankElement | TextElement | TagElement | ComponentElement


    export default function<S>(target?: HTMLElement, global?: Window): App<S>
}
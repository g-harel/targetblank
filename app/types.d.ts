declare module "targetblank" {
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
}

declare module "okwolo/lite" {
    export interface App<S> {
        (f: () => (state: S) => Element): void
        <P extends object>(path: string | RegExp, f: (params: P) => (state: S) => Element): void

        setState(state: NonNullable<S>): void
        setState(updater: (state: S) => NonNullable<S>): void

        getState(): S

        use(blob: "target", target: HTMLElement): void
    }

    export type BlankElement = boolean | null;

    export type TextElement = string | number;

    export type TagElement = {
        0: string;
        1?: object;
        2?: Element[];
    } & any[]

    export type ComponentElement<P extends object = {}, A extends object = {}> = {
        0: Component<P, A>;
        1?: P;
        2?: Element[];
    } & any[]

    export type Component<P extends object = {}, A extends object = {}> = (params: P) => (arg: A | undefined) => Element

    export type Element = BlankElement | TextElement | TagElement | ComponentElement

    export default function<S>(target?: HTMLElement, global?: Window): App<S>
}
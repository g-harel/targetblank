export type BlankElement = boolean | null;

export type TextElement = string | number;

export type TagElement = {
    0: string;
    1?: object;
    2?: Element[];
} & any[];

export type ComponentElement<P extends object = {}> = {
    0: Component<P>;
    1?: P;
    2?: Element[];
} & any[];

export type Component<P extends object = {}, D = undefined> = (
    params: P,
    update: (data?: D) => void,
) => any; // (data?: D) => Element;

export type Element =
    | BlankElement
    | TextElement
    | TagElement
    | ComponentElement;

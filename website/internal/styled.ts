import {style, types} from "typestyle";

export const styled = (tag: string) => (s: types.NestedCSSProperties) => {
    return `${tag} .${style(s)}` as any;
};

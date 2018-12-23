import {create} from "jss";
import {Style} from "jss/css";
import preset from "jss-preset-default";

const jss = create(preset());

export const styled = (tag: string) => (s: Style): string => {
    const name = "s";
    const {classes} = jss.createStyleSheet({[name]: s}).attach();
    return `${tag} .${classes[name]}`;
};

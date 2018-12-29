import {create} from "jss";
import {Style} from "jss/css";
import jssPresetDefault from "jss-preset-default";

const jss = create(jssPresetDefault());

export const styled = (tag: string) => (s: Style | Record<string, Style>) => {
    const name = "s";
    const {classes} = jss.createStyleSheet({[name]: s}).attach();
    return `${tag} .${classes[name]}` as any;
};

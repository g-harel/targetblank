import {create} from "jss";
import {Style} from "jss/css";
import preset from "jss-preset-default";

const jss = create(preset());

// TODO styled("div")
// TODO rate limiting (calling in render fn)
export const css = (s: Style): string => {
    const {classes} = jss.createStyleSheet({s}).attach();
    return classes[Object.keys(classes)[0] || ""];
};

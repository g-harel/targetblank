import {Style} from "jss/css";

export const reset: Style & Record<string, any> = {
    "-moz-osx-font-smoothing": "inherit",
    "-webkit-font-smoothing": "inherit",
    "-webkit-appearance": "none",
    background: "transparent",
    border: "none",
    color: "inherit",
    font: "inherit",
    lineHeight: "normal",
    margin: "0",
    outline: "none",
    overflow: "visible",
    padding: "0",
    textAlign: "inherit",
    width: "auto",

    "&::-moz-focus-inner": {
        padding: "0",
        border: "0",
    },

    "&:focus": {
        outline: "none",
    },
};

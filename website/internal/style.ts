import {style, types} from "typestyle";

// Helper to conveniently use typestyle alongside components.
export const styled = (tag: string) => (s: types.NestedCSSProperties) => {
    return `${tag} .${style(s)}` as any;
};

const unique = (s: types.NestedCSSProperties) => {
    s.$unique = true;
    return s;
};

// Helper to style element's placeholder.
export const placeholder = (s: types.NestedCSSProperties) => ({
    "&::placeholder": unique(s),
    "&::-webkit-input-placeholder": unique(s),
    "&::-moz-placeholder": unique(s),
    "&:-moz-placeholder": unique(s),
    "&:-ms-input-placeholder": unique(s),
});

// Button reset styles.
export const reset: types.NestedCSSProperties = {
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

    $nest: {
        "&::-moz-focus-inner": {
            padding: "0",
            border: "0",
        },

        "&:focus": {
            outline: "none",
        },
    },
};

// Standardized breakpoints.
export const breakpoint = {
    sm: "@media(max-width: 768px)",
};

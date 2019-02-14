import {style, types} from "typestyle";

// Helper to conveniently use typestyle alongside components.
export const styled = (tag: string) => (s: types.NestedCSSProperties) => {
    return `${tag} .${style(s)}` as any;
};

// Standardized breakpoints.
export const breakpoint = {
    sm: "@media(max-width: 768px)",
};

// Standardized font families.
export const fonts = {
    monospace:
        "SFMono-Regular, Consolas, Liberation Mono, Menlo, Courier, monospace",
    normal:
        // tslint:disable-next-line:quotemark
        '-apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, Oxygen, Ubuntu, ' +
        // tslint:disable-next-line:quotemark
        'Cantarell, "Open Sans", "Helvetica Neue", sans-serif',
};

// Standardized font sizes.
export const size = {
    tiny: "0.8rem",
    normal: "1rem",
    large: "1.2rem",
    title: "1.7rem",
};

// Standardized colors.
export const colors = {
    success: "yellowgreen",
    error: "tomato",
    decoration: "#e8e3e6",
    backgroundPrimary: "#ffffff",
    backgroundSecondary: "#faf8f6",
    textPrimary: "#332832",
    textSecondarySmall: "#766873",
    textSecondaryLarge: "#948c91",
};

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

// Helper to style element's placeholder.
export const placeholder = (s: types.NestedCSSProperties) => ({
    "&::placeholder": unique(s),
    "&::-webkit-input-placeholder": unique(s),
    "&::-moz-placeholder": unique(s),
    "&:-moz-placeholder": unique(s),
    "&:-ms-input-placeholder": unique(s),
});

const unique = (s: types.NestedCSSProperties) => {
    s.$unique = true;
    return s;
};

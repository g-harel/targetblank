// Detect if the os/browser is using a dark theme.
// This is the equivalent of the css media query.
const darkMode =
    window.matchMedia &&
    window.matchMedia("(prefers-color-scheme: dark)").matches;

// Standardized breakpoints.
export const breakpoint = {
    sm: "@media(max-width: 720px)",
    xs: "@media(max-width: 480px)",
};

// Standardized font families.
export const font = {
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
// prettier-ignore
export const color = darkMode ? {
    success: "yellowgreen",
    error: "tomato",
    decoration: "#363636",
    backgroundPrimary: "#282828",
    backgroundSecondary: "#1f1f1f",
    textPrimary: "#ffffff",
    textSecondarySmall: "#aaaaaa",
    textSecondaryLarge: "#aaaaaa",
} : {
    success: "yellowgreen",
    error: "tomato",
    decoration: "#e8e3e6",
    backgroundPrimary: "#ffffff",
    backgroundSecondary: "#faf8f6",
    textPrimary: "#332832",
    textSecondarySmall: "#766873",
    textSecondaryLarge: "#948c91",
};

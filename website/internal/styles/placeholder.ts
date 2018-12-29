import {Style} from "jss/css";

// Helper to style element's placeholder.
export const placeholder = (s: Style) => ({
    "&::-webkit-input-placeholder": s,
    "&::-moz-placeholder": s,
    "&:-moz-placeholder": s,
    "&:-ms-input-placeholder": s,
});

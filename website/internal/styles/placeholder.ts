import {types} from "typestyle";

// Helper to style element's placeholder.
export const placeholder = (s: types.NestedCSSProperties) => ({
    "&::-webkit-input-placeholder": s,
    "&::-moz-placeholder": s,
    "&:-moz-placeholder": s,
    "&:-ms-input-placeholder": s,
});

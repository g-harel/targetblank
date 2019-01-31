import {types} from "typestyle";

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

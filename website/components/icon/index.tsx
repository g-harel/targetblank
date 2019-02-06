import {keyframes} from "typestyle";

import {Component} from "../../internal/types";
import {styled} from "../../internal/style";

const spin = keyframes({
    "0%": {transform: "rotate(0deg)"},
    "100%": {transform: "rotate(359deg)"},
});

const Wrapper = styled("span")({
    display: "inline-block",
    fontSize: "1.1em",
    verticalAlign: "-0.17em",
    $nest: {
        "& svg": {
            fill: "currentColor",
        },

        "&.spin": {
            animationName: spin,
            animationDuration: "1s",
            animationIterationCount: "infinite",
            animationTimingFunction: "linear",
        },
    },
});

// tslint:disable:max-line-length
// https://github.com/FortAwesome/Font-Awesome-Pro
const paths: Record<Props["name"], string> = {
    check:
        "M256 8C119.033 8 8 119.033 8 256s111.033 248 248 248 248-111.033 248-248S392.967 8 256 8zm0 48c110.532 0 200 89.451 200 200 0 110.532-89.451 200-200 200-110.532 0-200-89.451-200-200 0-110.532 89.451-200 200-200m140.204 130.267l-22.536-22.718c-4.667-4.705-12.265-4.736-16.97-.068L215.346 303.697l-59.792-60.277c-4.667-4.705-12.265-4.736-16.97-.069l-22.719 22.536c-4.705 4.667-4.736 12.265-.068 16.971l90.781 91.516c4.667 4.705 12.265 4.736 16.97.068l172.589-171.204c4.704-4.668 4.734-12.266.067-16.971z",
    arrow:
        "M218.101 38.101L198.302 57.9c-4.686 4.686-4.686 12.284 0 16.971L353.432 230H12c-6.627 0-12 5.373-12 12v28c0 6.627 5.373 12 12 12h341.432l-155.13 155.13c-4.686 4.686-4.686 12.284 0 16.971l19.799 19.799c4.686 4.686 12.284 4.686 16.971 0l209.414-209.414c4.686-4.686 4.686-12.284 0-16.971L235.071 38.101c-4.686-4.687-12.284-4.687-16.97 0z",
    spinner:
        "M456.433 371.72l-27.79-16.045c-7.192-4.152-10.052-13.136-6.487-20.636 25.82-54.328 23.566-118.602-6.768-171.03-30.265-52.529-84.802-86.621-144.76-91.424C262.35 71.922 256 64.953 256 56.649V24.56c0-9.31 7.916-16.609 17.204-15.96 81.795 5.717 156.412 51.902 197.611 123.408 41.301 71.385 43.99 159.096 8.042 232.792-4.082 8.369-14.361 11.575-22.424 6.92z",
};

export interface Props {
    name: "check" | "arrow" | "spinner";
    color?: string;
    spin?: boolean;
    size?: number;
}

export const Icon: Component<Props> = (props) => () => (
    <Wrapper
        className={{spin: props.spin}}
        style={`color: ${props.color || "inherit"};`}
        innerHTML={`
            <svg
                xmlns="http://www.w3.org/2000/svg"
                width="${props.size || 1}em"
                height="${props.size || 1}em"
                viewBox="0 0 512 512"
            >
                <path d="${paths[props.name]}" />
            </svg>
        `}
    />
);

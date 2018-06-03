import "../static/component.input.scss";

import {Component} from "okwolo/lite";

export type props = {
    validator: (s: String) => string | null,
    placeholder: string,
};

export const input: Component<props> = (props) => {
    return () => (
        ["div.input", {}, [
            ["input", {
                type: "text",
                placeholder: " " + props.placeholder,
            }],
        ]]
    );
};

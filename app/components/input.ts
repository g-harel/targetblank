import "../static/component.input.scss";

import {Component} from "okwolo/lite";

export type props = {
    title?: string,
    callback?: (boolean) => void;
    validator: RegExp,
    message: string,
    placeholder: string,
};

export type args = {
    error: string;
};

export const input: Component<props, args> = (props, update) => {
    let timeout;

    const oninput = (event) => {
        const {value} = event.target;

        clearTimeout(timeout);

        if (value.length === 0) {
            return update({error: ""});
        }

        const valid = value.match(props.validator);

        if (props.callback) {
            props.callback(valid);
        }

        const error = valid ? "" : props.message;
        timeout = setTimeout(() => {
            update({error});
        }, 500);
    };

    return (args = {error: ""}) => (
        ["div.input", {}, [
            props.title ? ["span.title", {}, [
                props.title,
            ]] : "",
            ["input", {
                oninput,
                type: "text",
                placeholder: " " + props.placeholder,
            }],
            ["div.icon", {
                style: "color: #cfcfcf;",
            }, [
                ["i.far.fa-xs.fa-arrow-right"],
            ]],
            ["div.error", {}, [
                args.error,
            ]],
        ]]
    );
};

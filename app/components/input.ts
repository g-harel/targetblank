import "../static/component.input.scss";

import {Component} from "okwolo/lite";

export type props = {
    title?: string,
    callback?: (string) => Promise<string>;
    validator: RegExp,
    message: string,
    placeholder: string,
};

export const input: Component<props> = (props, update) => {
    let timeout;

    let error = "sad";
    let loading = false;
    let valid = false;
    let value = "";

    const oninput = (event) => {
        // reset any pending error message
        clearTimeout(timeout);

        value = event.target.value.trim();
        valid = !!value.match(props.validator);

        // immediately show new valid state
        update();

        // empty value does not show error (but is not valid)
        if (valid || value.length === 0) {
            error = "";
            update();
            return;
        }

        // delay error message
        timeout = setTimeout(() => {
            error = props.message;
            update();
        }, 750);
    };

    const onsubmit = async (e) => {
        e.preventDefault();

        if (!valid) {
            return;
        }

        loading = true;
        update();

        // show callback's error
        error = await props.callback(value) || "";

        // reset internal state
        loading = false;
        valid = false;
        value = "";

        update();
    };

    return () => (
        ["form.input", {
            onsubmit,
            className: {
                loading,
            },
        }, [
            props.title ? ["span.title", {}, [
                props.title,
            ]] : "",
            ["input", {
                value,
                oninput,
                type: "text",
                placeholder: " " + props.placeholder,
            }],
            ["button", {
                className: {
                    enabled: valid,
                },
                type: "submit",
            }, [
                loading
                    ? ["i.far.fa-xs.fa-spinner-third.fa-spin"]
                    : ["i.far.fa-xs.fa-arrow-right"],
            ]],
            ["div.error", {}, [
                error,
            ]],
        ]]
    );
};

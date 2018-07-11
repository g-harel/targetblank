import "../static/component.input.scss";

export type props = {
    title?: string,
    type?: string,
    callback?: (string) => Promise<string>;
    validator: RegExp,
    message: string,
    placeholder: string,
    focus: boolean,
};

const focusOnInput = () => {
    setTimeout(() => requestAnimationFrame(() => {
        const input: HTMLElement = document.querySelector("form.input input");
        if (input) {
            input.focus();
        }
    }));
};

export const input = (props: props, update) => {
    let error = "";
    let loading = false;
    let valid = false;
    let value = "";

    if (props.focus) {
        focusOnInput();
    }

    let timeout;

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

        focusOnInput();
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
                type: props.type || "text",
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

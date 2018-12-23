import "../../static/component.input.scss";

export interface Props {
    title?: string;
    type?: string;
    callback?: (string) => Promise<string>;
    validator: RegExp;
    message: string;
    placeholder: string;
    focus: boolean;
}

const focusOnInput = () => {
    setTimeout(() => requestAnimationFrame(() => {
        const input: HTMLElement = document.querySelector("form.input input");
        if (input) {
            input.focus();
        }
    }));
};

export const Input = (props: Props, update) => {
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
        <form
            className={{loading, input: true}}
            onsubmit={onsubmit}
        >
            {props.title ? (
                <span className="title">
                    {props.title}
                </span>
            ) : ""}
            <input
                type={props.type || "text"}
                value={value}
                oninput={oninput}
                placeholder={` ${props.placeholder}`}
            />
            <button type="submit" className={{enabled: valid}}>
                {loading ? (
                    <i className="far fa-xs fa-spinner-third fa-spin" />
                ) : (
                    <i className="far fa-xs fa-arrow-right" />
                )}
            </button>
            <div className="error">
                {error}
            </div>
        </form>
    );
};

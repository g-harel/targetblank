import {styled} from "../../internal/styled";
import {placeholder} from "../../internal/styles/placeholder";
import {reset} from "../../internal/styles/button";
import {Component} from "../../internal/types";

const width = 16;

const Form = styled("form")({
    width: `${width + 1}rem`,
    overflowX: "hidden",
    textAlign: "left",
    margin: "0 auto",

    $nest: {
        "&.loading": {
            filter: "grayscale(100%)",
            pointerEvents: "none",

            $nest: {
                "& *:not(button)": {
                    opacity: 0.5,
                },

                "& button": {
                    color: "#555",
                    cursor: "default",
                },
            },
        },
    },
});

const Title = styled("span")({
    display: "block",
    fontSize: "1.3rem",
    fontWeight: "bold",
    margin: "0.5rem 0.9rem 0.1rem",
    width: `${width}rem`,
});

const StyledInput = styled("input")({
    border: "1px solid #bbb",
    borderRadius: "2px",
    height: "1.85rem",
    margin: "0.3rem 0.5rem 0",
    outline: "0",
    padding: "0.25rem 1.8rem 0.25rem 0.5rem",
    width: `${width}rem`,

    $nest: {
        "&:focus": {
            boxShadow: "0 0 1px 1px #eee",
            borderColor: "inherit",
        },

        ...placeholder({
            color: "#cfcfcf",
            opacity: 1,
        }),
    },
});

const Button = styled("button")({
    ...reset,

    transform: "translate(-0.35rem, -1.95rem)",
    padding: "0.65rem 0.7rem 0.65rem",
    display: "inline-block",
    pointerEvents: "none",
    cursor: "default",
    color: "#cfcfcf",
    float: "right",

    $nest: {
        "&.enabled": {
            pointerEvents: "all",
            cursor: "pointer",
            color: "inherit",
        },
    },
});

const Error = styled("div")({
    color: "tomato",
    fontSize: "0.75rem",
    height: "1.15rem",
    margin: "0.35rem 1rem",
    width: `${width}rem`,
});

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
    setTimeout(() =>
        requestAnimationFrame(() => {
            const input: HTMLElement = document.querySelector("form input");
            if (input) {
                input.focus();
            }
        }),
    );
};

export const Input: Component<Props> = (props, update) => {
    let error = "";
    let loading = false;
    let valid = false;
    let value = "";

    if (props.focus) {
        focusOnInput();
    }

    let timeout;

    const oninput = (event) => {
        // Reset any pending error message.
        clearTimeout(timeout);

        value = event.target.value.trim();
        valid = !!value.match(props.validator);

        // Immediately show new valid state.
        update();

        // Empty value does not show error (but is not valid).
        if (valid || value.length === 0) {
            error = "";
            update();
            return;
        }

        // Delay error message until typing stops.
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

        // Callback's error is displayed.
        error = (await props.callback(value)) || "";

        // Reset internal state after submit.
        loading = false;
        valid = false;
        value = "";

        update();

        focusOnInput();
    };

    return () => (
        <Form className={{loading}} onsubmit={onsubmit}>
            {!!props.title && <Title>{props.title}</Title>}
            <StyledInput
                type={props.type || "text"}
                value={value}
                oninput={oninput}
                placeholder={` ${props.placeholder}`}
            />
            <Button type="submit" className={{enabled: valid}}>
                {loading ? (
                    <i className="far fa-xs fa-spinner-third fa-spin" />
                ) : (
                    <i className="far fa-xs fa-arrow-right" />
                )}
            </Button>
            <Error>{error}</Error>
        </Form>
    );
};

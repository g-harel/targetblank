import {styled, reset, placeholder, colors, size} from "../../internal/style";
import {Component} from "../../internal/types";
import {Icon} from "../icon";

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
                    opacity: 0.7,
                },

                "& button": {
                    color: colors.textPrimary,
                    cursor: "default",
                },
            },
        },
    },
});

const Title = styled("span")({
    display: "flex",
    fontWeight: 600,
    alignItems: "center",
    margin: "0.5rem 0.9rem 0.1rem",
    width: `${width}rem`,

    $nest: {
        "&.error": {
            color: colors.error,
        },
    },
});

const Hint = styled("span")({
    color: colors.textSecondaryLarge,
    fontSize: size.tiny,
    marginLeft: "0.5rem",

    $nest: {
        "&.error": {
            color: colors.error,
        },
    },
});

const StyledInput = styled("input")({
    border: `1px solid ${colors.textSecondarySmall}`,
    borderRadius: "2px",
    boxShadow: "none",
    height: "1.85rem",
    margin: "0.3rem 0.5rem 0",
    padding: "1rem 1.8rem 1rem 0.5rem",
    width: `${width}rem`,

    $nest: {
        "&:focus": {
            boxShadow: `0 0 0.5px 1px ${colors.decoration}`,
            borderColor: colors.textPrimary,
        },

        "&.error": {
            borderColor: colors.error,
        },

        ...placeholder({
            color: colors.textSecondarySmall,
            opacity: 1,
        }),
    },
});

const Button = styled("button")({
    ...reset,

    transform: "translate(-0.5rem, -2.1rem)",
    padding: "0.45rem 0.7rem",
    display: "inline-block",
    pointerEvents: "none",
    cursor: "default",
    color: colors.textSecondaryLarge,
    float: "right",
    lineHeight: 1,

    $nest: {
        "&.enabled": {
            pointerEvents: "all",
            cursor: "pointer",
            color: "inherit",
        },

        "&.error": {
            color: colors.error,
        },
    },
});

const Error = styled("div")({
    color: colors.error,
    fontSize: size.tiny,
    fontWeight: 600,
    height: "1.15rem",
    margin: "0.35rem 1rem",
    width: `${width}rem`,
});

export interface Props {
    title?: string;
    hint?: string;
    type?: string;
    autocomplete?: string;
    callback: (value: string) => Promise<string>;
    validator: RegExp;
    message: string;
    placeholder: string;
    focus: boolean;
}

const focusOnInput = () => {
    setTimeout(() =>
        requestAnimationFrame(() => {
            const input = document.querySelector<HTMLElement>("form input");
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

    let timeout: any;

    const oninput = (event: TextEvent) => {
        // Reset any pending error message.
        clearTimeout(timeout);

        value = (event.target as HTMLInputElement).value.trim();
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

    const onsubmit = async (e: KeyboardEvent) => {
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
            <Title className={{error}}>
                {!!props.title && props.title}
                {!!props.hint && <Hint className={{error}}>{props.hint}</Hint>}
            </Title>
            <StyledInput
                type={props.type || "text"}
                value={value}
                oninput={oninput}
                placeholder={`${props.placeholder}`}
                className={{error}}
            />
            <Button type="submit" className={{error, enabled: valid}}>
                {loading ? (
                    <Icon name="spinner" spin size={0.8} />
                ) : (
                    <Icon name="arrow" size={0.8} />
                )}
            </Button>
            <Error>{error}</Error>
        </Form>
    );
};

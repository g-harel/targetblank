import {styled} from "../../internal/styled";
import {breakpoint} from "../../internal/styles/breakpoint";
import {Component} from "../../internal/types";

const Wrapper = styled("div")({
    backgroundColor: "#f88",
    border: "1px solid #800",
    borderBottom: "none",
    borderRadius: "2px 2px 0 0",
    bottom: "0",
    color: "#a00",
    position: "fixed",
    right: "100px",
    width: "400px",

    "&.hidden": {
        display: "none",
    },

    [breakpoint.sm]: {
        borderLeft: "none",
        borderRigth: "none",
        right: "0",
        width: "100%",
    },
});

const Title = styled("div")({
    borderBottom: "1px solid #a00",
    fontWeight: "bold",
    letterSpacing: "1px",
    padding: "6px 12px",
});

const Dismiss = styled("div")({
    cursor: "pointer",
    float: "right",
    fontSize: "0.8em",
    fontWeight: "normal",
    padding: "5px",
    transform: "translate(2px, -2px)",
    userSelect: "none",
});

const Error = styled("div")({
    fontFamily: "'Inconsolata', monospace",
    fontSize: "0.8em",
    padding: "0 12px",
    margin: "12px 0",
});

export interface Props {
    errors: string[];
    hide: () => void;
}

export const Errors: Component<Props> = ({errors, hide}) => () => (
    <Wrapper className={{
        hidden: !errors.length,
    }}>
        <Title>
            error
            <Dismiss onclick={hide}>
                dismiss
            </Dismiss>
        </Title>
        {...errors.map((err) => (
            <Error>
                {err}
            </Error>
        ))}
    </Wrapper>
);

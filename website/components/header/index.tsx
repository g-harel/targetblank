import {styled} from "../../internal/style";
import {Component} from "../../internal/types";

const Wrapper = styled("header")({
    "-moz-user-select": "-moz-none",
    fontFamily: "Arial, Helvetica Neue, Helvetica, sans-serif",
    fontSize: "1.7rem",
    fontWeight: 600,
    lineHeight: "2.9rem",
    padding: "4.5rem 1.4rem 4rem",
    textAlign: "center",
    userSelect: "none",
    pointerEvents: "none",
    width: "100%",

    $nest: {
        "&.muted": {
            opacity: 0.4,
        },
    },
});

export interface Props {
    title?: string;
    muted?: boolean;
}

export const Header: Component<Props> = (props) => () => (
    <Wrapper className={{muted: props.muted}}>
        {props.title === undefined ? "targetblank" : props.title}
        &nbsp;
    </Wrapper>
);

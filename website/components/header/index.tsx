import {styled} from "../../internal/style";
import {size, color} from "../../internal/style/theme";
import {Component} from "../../internal/types";

const Wrapper = styled("header")({
    "-moz-user-select": "-moz-none",
    fontFamily: "Arial, Helvetica Neue, Helvetica, sans-serif",
    fontSize: size.title,
    fontWeight: 600,
    lineHeight: "2.9rem",
    padding: "4.5rem 1.4rem 4rem",
    textAlign: "center",
    userSelect: "none",
    pointerEvents: "none",
    width: "100%",

    $nest: {
        "&.muted": {
            color: color.textSecondaryLarge,
        },
    },
});

const Subtitle = styled("div")({
    color: color.textSecondarySmall,
    height: 0,
    fontSize: size.normal,
    transform: "translateY(-1rem)",
});

export interface Props {
    title?: string;
    muted?: boolean;
    sub?: string;
}

export const Header: Component<Props> = (props) => () =>
    (
        <Wrapper className={{muted: props.muted}}>
            &nbsp;
            {props.title === undefined ? "targetblank" : props.title}
            &nbsp;
            {props.sub ? <Subtitle>{props.sub}</Subtitle> : ""}
        </Wrapper>
    );

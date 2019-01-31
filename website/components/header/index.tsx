import {styled} from "../../internal/styled";
import {Component} from "../../internal/types";

const Wrapper = styled("header")({
    "-moz-user-select": "-moz-none",
    font: "800 1.6rem 'Lora', serif",
    lineHeight: "2.9rem",
    padding: "4.5rem 1.4rem 4rem",
    textAlign: "center",
    userSelect: "none",
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
        {props.title || "targetblank"}
    </Wrapper>
);

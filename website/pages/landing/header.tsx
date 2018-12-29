import {styled} from "../../internal/styled";
import {breakpoint} from "../../internal/styles/breakpoint";
import {Component} from "../../internal/types";

export const headerHeight = 2.9;

const Wrapper = styled("div")({
    borderBottom: "1px solid #ddd",
    font: "800 1.6rem 'Lora', serif",
    height: `${headerHeight}rem`,
    padding: "0.3rem 1.4rem 0.5rem",
    userSelect: "none",

    [breakpoint.sm]: {
        textAlign: "center",
    },
});

export const Header: Component = () => () => <Wrapper>targetblank</Wrapper>;

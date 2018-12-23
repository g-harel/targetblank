import {styled} from "../../library/styled";
import {breakpoint} from "../../library/styles/breakpoint";
import {Component} from "../../library/types";

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

export const Header: Component = () => () => (
    <Wrapper>
        targetblank
    </Wrapper>
);

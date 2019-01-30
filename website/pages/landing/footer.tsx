import {styled} from "../../internal/styled";
import {Component} from "../../internal/types";
import {Anchor} from "../../components/anchor";

export const footerHeight = 2.8;

const Wrapper = styled("footer")({
    height: `${footerHeight}rem`,
    overflow: "hidden",
    textAlign: "center",
    width: "100vw",
});

const Icon = styled("i")({
    color: "#eee",
    fontSize: "1.6rem",
    transition: "all 0.2s ease",

    "&:hover": {
        color: "#ddd",
    },
});

export const Footer: Component = () => () => (
    <Wrapper>
        <Anchor href="https://github.com/g-harel/targetblank">
            <Icon className="fab fa-github" />
        </Anchor>
    </Wrapper>
);

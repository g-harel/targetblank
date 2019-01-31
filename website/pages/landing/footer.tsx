import {styled} from "../../internal/styled";
import {Component} from "../../internal/types";
import {Anchor} from "../../components/anchor";

const Wrapper = styled("footer")({
    bottom: 0,
    overflow: "hidden",
    padding: "1rem",
    position: "absolute",
    textAlign: "center",
    width: "100vw",
});

const Icon = styled("i")({
    color: "#eee",
    fontSize: "1.6rem",
    transition: "all 0.2s ease",

    $nest: {
        "&:hover": {
            color: "#ddd",
        },
    },
});

export const Footer: Component = () => () => (
    <Wrapper>
        <Anchor href="https://github.com/g-harel/targetblank">
            <Icon className="fab fa-github" />
        </Anchor>
    </Wrapper>
);

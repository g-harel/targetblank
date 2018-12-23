import {styled} from "../../library/styled";

export const footerHeight = 2.8;

const Wrapper = styled("div")({
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

export const Footer = () => () => (
    <Wrapper>
        <a href="https://github.com/g-harel/targetblank">
            <Icon className="fab fa-github" />
        </a>
    </Wrapper>
);

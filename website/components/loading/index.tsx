import {Component} from "../../internal/types";
import {styled} from "../../internal/styled";

const Wrapper = styled("div")({
    backgroundColor: "#fafafa",
    font: "800 1.6rem 'Lora', serif",
    height: "100%",
    paddingTop: "22.5vh",
    textAlign: "center",
    opacity: 0.2,
    userSelect: "none",
    width: "100%",
});

export const Loading: Component = () => () => {
    return (
        <Wrapper>
            targetblank
        </Wrapper>
    );
};

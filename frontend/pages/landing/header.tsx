import {styled} from "../../library/styled";

export const headerHeight = 2.9;

const Wrapper = styled("div")({
    borderBottom: "1px solid #ddd",
    font: "800 1.6rem 'Lora', serif",
    height: `${headerHeight}rem`,
    padding: "0.3rem 1.4rem 0.5rem",
    userSelect: "none",

    "@media (max-width: 768px)": {
        textAlign: "center",
    },
});

export const Header = () => () => (
    <Wrapper>
        targetblank
    </Wrapper>
);

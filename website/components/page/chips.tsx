import {styled} from "../../internal/style";
import {size, color} from "../../internal/style/theme";
import {Component} from "../../internal/types";

const chips: Array<{
    text: string;
}> = [];

const Wrapper = styled("div")({
    color: color.textSecondarySmall,
    height: 0,
    fontSize: size.normal,
});

const Chip = styled("div")({
    border: "1px solid red",
    margin: "1rem",
    padding: "1rem",
});

export interface Props {}

export const Chips: Component<Props> = () => () => (
    <Wrapper>
        {...chips.map((chip) => (
            <Chip>{chip.text}</Chip>
        ))}
    </Wrapper>
);

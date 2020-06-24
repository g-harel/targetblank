import {styled} from "../../internal/style";
import {size, color} from "../../internal/style/theme";
import {Component} from "../../internal/types";

const chips: Array<{
    text: string;
}> = [
    // {text: "You're already logged in!"},
    // {text: "This comment is multi-line because I need to check out what that would look like."},
    // {text: "This one shows how the chips look like when there are many of them."},
];

const Wrapper = styled("div")({
    alignItems: "center",
    display: "flex",
    flexDirection: "column",
    fontSize: size.normal,
    height: 0,
    zIndex: 1,
});

const Chip = styled("div")({
    backgroundColor: color.backgroundPrimary,
    border: "1px solid transparent",
    borderRadius: "2px",
    boxShadow: `0 0 0.5px 1px ${color.decoration}`,
    color: color.textSecondarySmall,
    fontWeight: 600,
    marginTop: "1rem",
    padding: "0.5rem 0.5rem 0.6rem",
    maxWidth: "30rem",
});

export interface Props {}

export const Chips: Component<Props> = () => () => (
    <Wrapper>{...chips.map((chip) => <Chip>{chip.text}</Chip>)}</Wrapper>
);

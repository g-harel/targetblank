import {keyframes} from "typestyle";

import {styled} from "../../internal/style";
import {size, color} from "../../internal/style/theme";
import {Component} from "../../internal/types";

// Current state of the chip.
// Shared between renders of the component.
let chip: null | {
    text: string;
    timeout: any;
} = null;

// Cleat the current chip.
const clearChip = () => {
    chip = null;
    updateComponent();
};

// Show the given text in the chip for the provided amount of time.
export const showChip = (text: string, ttl: number) => {
    if (chip != null) {
        clearTimeout(chip.timeout);
    }
    // Clear chip before recreating so that a new node is animated in.
    clearChip();
    chip = {
        text: text,
        timeout: setTimeout(clearChip, ttl),
    };
    updateComponent();
};

// Placeholder to contain the most recent Chip component update function.
// This restricts the component to only function when there is a single instance of it.
let updateComponent: Function = () => {};

const fadeIn = keyframes({
    "0%": {opacity: 0},
    "100%": {opacity: 1},
});

const Wrapper = styled("div")({
    alignItems: "center",
    display: "flex",
    flexDirection: "column",
    fontSize: size.normal,
    height: 0,
    zIndex: 999,
});

const Chip = styled("div")({
    animation: `${fadeIn} 0.2s ease`,
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

export const Chips: Component<Props> = (_, update) => {
    updateComponent = update;
    return () =>
        chip && (
            <Wrapper>
                <Chip>{chip.text}</Chip>
            </Wrapper>
        );
};

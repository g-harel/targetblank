import {Component, IPageEntry} from "../../internal/types";
import {styled} from "../../internal/style";
import {color} from "../../internal/style/theme";
import {Anchor} from "../../components/anchor";
import {Icon} from "../../components/icon";

const Wrapper = styled("div")({
    color: color.textPrimary,
    fontWeight: 600,
    lineHeight: "1.6rem",
    padding: "0 0.3rem",

    $nest: {
        "& &": {
            paddingLeft: "2rem",
        },

        "& > span, & > a": {
            display: "inline-block",
            lineHeight: 1.2,
            margin: "0.3rem 0 0",
            padding: "0.3em 0.4em",
        },

        "& > a:focus": {
            backgroundColor: color.decoration,
            borderRadius: "4px",
            outline: "none",
        },
    },
});

const Indicator = styled("div")({
    display: "inline-block",
    height: 0,
    padding: 0,
    transform: "translate(-1.1em, -0.15em)",
    width: 0,
});

const ItemTitle = styled("span")({
    color: color.textSecondarySmall,
});

export interface Props {
    item: IPageEntry;
    generateID: () => string;
    isHighlighted: (item: IPageEntry, key: string) => boolean;
}

export const Item: Component<Props> = (props) => () => {
    const id = props.generateID();
    const isHighlighted = props.isHighlighted(props.item, id);

    let Title = null;
    if (props.item.link) {
        Title = (
            <Anchor id={id} href={props.item.link}>
                {props.item.label}
            </Anchor>
        );
    } else {
        Title = <ItemTitle>{props.item.label}</ItemTitle>;
    }

    return (
        <Wrapper>
            {isHighlighted && (
                <Indicator>
                    <Icon
                        name="arrow"
                        size={0.7}
                        color={color.textSecondaryLarge}
                    />
                </Indicator>
            )}
            {Title}
            {...props.item.entries.map((item) => (
                <Item
                    item={item}
                    isHighlighted={props.isHighlighted}
                    generateID={props.generateID}
                />
            ))}
        </Wrapper>
    );
};

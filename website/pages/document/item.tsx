import {Component, IPageEntry} from "../../internal/types";
import {styled} from "../../internal/style";
import {Anchor} from "../../components/anchor";

const Wrapper = styled("div")({
    color: "#333",
    fontWeight: 600,
    lineHeight: "1.6rem",
    padding: "0.6rem 0.3rem 0",

    $nest: {
        "& &": {
            paddingLeft: "2rem",
        },

        "&.highlighted": {
            color: "yellowgreen",
        },
    },
});

const ItemTitle = styled("span")({
    color: "#888",
});

export interface Props {
    item: IPageEntry;
    isHighlighted: (item: IPageEntry) => boolean;
}

export const Item: Component<Props> = (props) => () => {
    let Title = null;
    if (props.item.link) {
        Title = <Anchor href={props.item.link}>{props.item.label}</Anchor>;
    } else {
        Title = <ItemTitle>{props.item.label}</ItemTitle>;
    }

    return (
        <Wrapper className={{highlighted: props.isHighlighted(props.item)}}>
            {Title}
            {...props.item.entries.map((item) => (
                <Item item={item} isHighlighted={props.isHighlighted} />
            ))}
        </Wrapper>
    );
};

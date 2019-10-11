import {Component, IPageEntry} from "../../internal/types";
import {styled} from "../../internal/style";
import {color} from "../../internal/style/theme";
import {Anchor} from "../../components/anchor";

const Wrapper = styled("div")({
    color: color.textPrimary,
    fontWeight: 600,
    lineHeight: "1.6rem",
    padding: "0.6rem 0.3rem 0",

    $nest: {
        "& &": {
            paddingLeft: "2rem",
        },

        "&.highlighted": {
            color: color.error,
        },
    },
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
        <Wrapper className={{highlighted: props.isHighlighted(props.item, id)}}>
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

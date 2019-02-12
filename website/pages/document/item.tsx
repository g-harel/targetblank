import {Component, IPageItem} from "../../internal/types";
import {styled} from "../../internal/style";
import {Anchor} from "../../components/anchor";

const Wrapper = styled("div")({
    fontWeight: 600,
    lineHeight: "1.6rem",
    padding: "0.6rem 0.3rem 0",

    $nest: {
        "& &": {
            paddingLeft: "2rem",
        },
    },
});

const ItemTitle = styled("span")({
    color: "#888",
});

export const Item: Component<IPageItem> = (props) => () => {
    let Title = null;
    if (props.link) {
        Title = <Anchor href={props.link}>{props.label}</Anchor>;
    } else {
        Title = <ItemTitle>{props.label}</ItemTitle>;
    }

    return (
        <Wrapper>
            {Title}
            {...props.entries.map((item) => <Item {...item} />)}
        </Wrapper>
    );
};

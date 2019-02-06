import {Component, IPageItem} from "../../internal/types";
import {styled} from "../../internal/style";
import {Anchor} from "../../components/anchor";

const Wrapper = styled("div")({
    fontWeight: 700,
    lineHeight: "1.6rem",
    padding: "0.3rem",

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

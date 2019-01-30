import {IPageItem} from "../../internal/client/types";
import {Component} from "../../internal/types";
import {styled} from "../../internal/styled";
import {Anchor} from "../../components/anchor";

const Wrapper = styled("div")({
    fontSize: "1.1rem",
    fontWeight: 700,
    lineHeight: "1.6rem",
    padding: "0.3rem",

    "& &": {
        paddingLeft: "2rem",
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

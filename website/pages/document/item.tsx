import {IPageItem} from "../../internal/client/types";
import {Component} from "../../internal/types";
import {styled} from "../../internal/styled";

const Wrapper = styled("div")({
    fontSize: "1.1rem",
    fontWeight: 700,
    lineHeight: "1.6rem",
    padding: "0.3rem",

    "& &": {
        paddingLeft: "2rem",
    },
});

const ItemLink = styled("a")({
    color: "#888",
});

const ItemTitle = styled("span")({});

export const Item: Component<IPageItem> = (props) => () => {
    const Title = props.link ? ItemLink : ItemTitle;
    return (
        <Wrapper>
            <Title href={props.link}>
                {props.label}
            </Title>
            {...props.entries.map((item) => <Item {...item} />)}
        </Wrapper>
    );
};

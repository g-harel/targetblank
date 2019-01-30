import {app} from "../../internal/app";
import {client, IPageData} from "../../internal/client";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";
import {Loading} from "../../components/loading";
import {Item} from "./item";
import {Anchor} from "../../components/anchor";

const Wrapper = styled("div")({
    height: "100%",
    maxWidth: "1200px",
    margin: "0 auto",
});

const Edit = styled("div")({
    display: "inline",
    float: "right",
    fontWeight: "bold",
    padding: "0.85rem 1.4rem",
    position: "absolute",
    right: 0,
    userSelect: "none",
});

const Groups = styled("div")({
    alignContent: "center",
    display: "flex",
    flexWrap: "wrap",
    justifyContent: "center",
    maxWidth: "100%",
    minHeight: "100%",
    padding: "2rem 1rem",
});

const Group = styled("div")({
    border: "1px solid #eee",
    borderRadius: "2px",
    flexBasis: "30%",
    flexGrow: 1,
    flexShrink: 0,
    margin: "1rem",
    padding: "1rem 1.4rem",
});

const Items = styled("div")({});

export const Document: PageComponent<IPageData> = ({addr}, update) => {
    client.page.fetch(update, () => app.redirect(`/${addr}/login`), addr);

    return (data: IPageData) => {
        // Response not yet received.
        if (!data) return <Loading />;

        return (
            <Wrapper>
                <Edit>
                    <Anchor href={`/${addr}/edit`}>edit</Anchor>
                </Edit>
                <Groups>
                    {...data.groups.map((group) => (
                        <Group>
                            <Items>
                                {...group.entries.map((item) => (
                                    <Item {...item} />
                                ))}
                            </Items>
                        </Group>
                    ))}
                </Groups>
            </Wrapper>
        );
    };
};

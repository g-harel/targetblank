import {app} from "../../internal/app";
import {client, IPageData} from "../../internal/client";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";
import {Loading} from "../../components/loading";
import {Item} from "./item";
import {Anchor} from "../../components/anchor";
import {Header} from "../../components/header";

const Wrapper = styled("div")({
    display: "flex",
    flexDirection: "column",
    height: "100%",
    maxWidth: "1200px",
    margin: "0 auto",
});

const Action = styled("div")({
    "-moz-user-select": "-moz-none",
    float: "right",
    fontWeight: "bold",
    padding: "0.85rem 1.4rem",
    position: "absolute",
    right: 0,
    userSelect: "none",
});

const Edit = styled("div")({
    $nest: {
        "&:hover": {
            $nest: {
                "&::before": {
                    content: `"ctrl + enter"`,
                    opacity: 0.3,
                    margin: "0 0.7rem",
                },
            },
        },
    },
});

const Groups = styled("div")({
    display: "flex",
    flexWrap: "wrap",
    justifyContent: "center",
    padding: "0 1rem 3rem",
});

const Group = styled("div")({
    border: "1px solid #eee",
    borderRadius: "2px",
    flexBasis: "30%",
    flexGrow: 0,
    flexShrink: 0,
    margin: "1rem",
    padding: "1rem 1.4rem",
});

const Items = styled("div")({});

export const Document: PageComponent<IPageData> = ({addr}, update) => {
    client.page.fetch(update, () => app.redirect(`/${addr}/login`), addr);

    // Navigate to the edit page with "ctrl+enter".
    window.onkeypress = (e) => {
        const ctrl = navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey;
        if (ctrl && e.key === "Enter") {
            app.redirect(`/${addr}/edit`);
        }
    };

    return (data: IPageData) => {
        // Response not yet received.
        if (!data) return <Loading />;

        return (
            <Wrapper>
                <Action>
                    {client.page.auth(addr) ? (
                        <Edit>
                            <Anchor href={`/${addr}/edit`}>edit</Anchor>
                        </Edit>
                    ) : (
                        <Anchor href={`/${addr}/login`}>login</Anchor>
                    )}
                </Action>
                <Header muted title={data.meta.title} />
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

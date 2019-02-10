import {client, IPageData} from "../../internal/client";
import {PageComponent} from "../../components/page";
import {styled, breakpoint} from "../../internal/style";
import {Loading} from "../../components/loading";
import {Item} from "./item";
import {Anchor} from "../../components/anchor";
import {Header} from "../../components/header";
import {keyboard} from "../../internal/keyboard";
import {path, routes, redirect} from "../../routes";

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
    fontWeight: 600,
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
    margin: "0 1rem 2rem",
    padding: "1rem 1.4rem",

    $nest: {
        [breakpoint.sm]: {
            flexGrow: 1,
        },
    },
});

const Items = styled("div")({});

export const Document: PageComponent<IPageData> = ({addr}, update) => {
    client(addr!).pageRead(update, () => redirect(routes.login, addr!));

    // Navigate to the edit page with "ctrl+enter".
    keyboard((e) => {
        if (e.ctrl && e.key === "Enter") {
            redirect(routes.edit, addr!);
        }
    });

    return (data: IPageData) => {
        // Response not yet received.
        if (!data) return <Loading />;

        document.title = data.meta.title || "targetblank";

        return (
            <Wrapper>
                <Action>
                    {client(addr!).isAuthorized() ? (
                        <Edit>
                            <Anchor href={path(routes.edit, addr!)}>
                                edit
                            </Anchor>
                        </Edit>
                    ) : (
                        <Anchor href={path(routes.login, addr!)}>login</Anchor>
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

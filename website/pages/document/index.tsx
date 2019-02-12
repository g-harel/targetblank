import {client, IPageData} from "../../internal/client";
import {PageComponent} from "../../components/page";
import {styled, breakpoint} from "../../internal/style";
import {Loading} from "../../components/loading";
import {Item, Props as ItemProps} from "./item";
import {Anchor} from "../../components/anchor";
import {Header} from "../../components/header";
import {keyboard} from "../../internal/keyboard";
import {path, routes, redirect} from "../../routes";

const keyboardTimeout = 1000;

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
    padding: "1rem 1.4rem 1.6rem",

    $nest: {
        [breakpoint.sm]: {
            flexGrow: 1,
        },
    },
});

const Items = styled("div")({});

export const Document: PageComponent = ({addr}, update) => {
    let highlight: string = "";
    let timer: any = null;
    let data: IPageData | null = null;

    client(addr!).pageRead(
        (d) => {
            data = d;
            update();
        },
        () => redirect(routes.login, addr!),
    );

    keyboard((e) => {
        // Navigate to the edit page with "ctrl+enter".
        if (e.ctrl && e.key === "Enter") {
            redirect(routes.edit, addr!);
            return;
        }

        // Update highlight on keboard clicks.
        if (e.key.match(/^[a-z]$/g)) {
            highlight += e.key;
            clearTimeout(timer);
            timer = setTimeout(() => {
                highlight = "";
                update();
            }, keyboardTimeout);
            update();
        }
    });

    return () => {
        // Response not yet received.
        if (!data) return <Loading />;

        document.title = data.meta.title || "targetblank";

        // Checker given to pick the highlighted item.
        let found = false;
        const isHighlighted: ItemProps["isHighlighted"] = (item) => {
            if (found) {
                return false;
            }
            if (highlight.length === 0) {
                return false;
            }
            if (!item.link || !item.label) {
                return false;
            }

            // Format string to lowercase and remove accents.
            const formattedString = item.label
                .toLowerCase()
                .normalize("NFD")
                .replace(/[\u0300-\u036f]/g, "");
            if (formattedString.indexOf(highlight) >= 0) {
                found = true;
                return true;
            }

            return false;
        };

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
                                    <Item
                                        item={item}
                                        isHighlighted={isHighlighted}
                                    />
                                ))}
                            </Items>
                        </Group>
                    ))}
                </Groups>
            </Wrapper>
        );
    };
};

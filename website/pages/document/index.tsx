import {client, IPageData} from "../../internal/client";
import {PageComponent} from "../../components/page";
import {styled, breakpoint, colors} from "../../internal/style";
import {Loading} from "../../components/loading";
import {Item, Props as ItemProps} from "./item";
import {Anchor} from "../../components/anchor";
import {Header} from "../../components/header";
import {keyboard} from "../../internal/keyboard";
import {path, routes, safeRedirect} from "../../routes";
import {IPageEntry} from "../../internal/types";

const keyboardTimeout = 1000;

const Wrapper = styled("div")({
    display: "flex",
    flexDirection: "column",
    height: "100%",
    maxWidth: "1200px",
    margin: "0 auto",
    width: "100%",
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
                    content: `"shift + e"`,
                    color: colors.textSecondaryLarge,
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
    padding: "0 1rem",
});

const Group = styled("div")({
    border: `1px solid ${colors.decoration}`,
    borderRadius: "2px",
    flexBasis: "30%",
    flexShrink: 0,
    flexGrow: 0,
    margin: "0 1rem 2rem",
    padding: "1rem 1.4rem 1.6rem",

    $nest: {
        [breakpoint.sm]: {
            flexBasis: "50%",
        },
    },
});

const Items = styled("div")({});

export const Document: PageComponent = ({addr}, update) => {
    let data: IPageData | null = null;

    client(addr!).pageRead(
        (d) => {
            data = d;
            update();
        },
        () => safeRedirect(routes.login, addr!),
    );

    let highlight: string = "";
    let highlighted: IPageEntry | null = null;
    let highlighTimer: any = null;

    keyboard((e) => {
        // Navigate to the edit page with keyboard.
        if (e.shift && e.key === "E") {
            safeRedirect(routes.edit, addr!);
            return;
        }

        // Update highlight on keyboard clicks.
        if (e.key.match(/^[a-z0-9 -]$/g)) {
            highlight += e.key;
            clearTimeout(highlighTimer);
            highlighTimer = setTimeout(() => {
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

        // Also give focus to most recently highlighted link.
        setTimeout(() => {
            if (!highlighted) return;
            const link = document.querySelector(
                `a[href="${highlighted.link}"]`,
            );
            if (link) (link as any).focus();
        });

        // Checker given to pick the highlighted item.
        highlighted = null;
        const isHighlighted: ItemProps["isHighlighted"] = (item) => {
            if (highlighted) {
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
                highlighted = item;
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

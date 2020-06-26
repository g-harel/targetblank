import {client, IPageData} from "../../internal/client";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/style";
import {color} from "../../internal/style/theme";
import {Loading} from "../../components/loading";
import {Anchor} from "../../components/anchor";
import {Editor} from "../../components/editor";
import {Icon} from "../../components/icon";
import {keyboard} from "../../internal/keyboard";
import {path, routes, safeRedirect} from "../../routes";
import {showChip} from "../../components/page/chips";

const headerHeight = "2.9rem";
const saveDelay = 1400;

const Wrapper = styled("div")({
    backgroundColor: color.backgroundSecondary,
    display: "flex",
    flexDirection: "column",
    flexGrow: 1,
    minHeight: "100%",
});

const Header = styled("header")({
    "-moz-user-select": "-moz-none",
    backgroundColor: color.backgroundPrimary,
    borderBottom: `1px solid ${color.decoration}`,
    height: headerHeight,
    padding: "0.85rem 1.4rem",
    position: "fixed",
    userSelect: "none",
    width: "100%",
    zIndex: 3,
});

// Inline element to ensure no content is rendered under the header.
const HeaderPlaceholder = styled("div")({
    height: headerHeight,
});

const Done = styled("div")({
    display: "inline-block",
    float: "right",
    fontWeight: 600,

    $nest: {
        "&.disabled": {
            pointerEvents: "none",
            color: color.textSecondaryLarge,
        },

        "&:hover": {
            $nest: {
                "&::before": {
                    content: `"esc"`,
                    color: color.textSecondaryLarge,
                    margin: "0 0.7rem",
                },
            },
        },
    },
});

const Status = styled("div")({
    display: "inline-block",
    fontWeight: 600,

    $nest: {
        "&.error": {
            color: color.error,
        },
    },
});

export interface Data {
    value: string;
    status: "saving" | "saved" | "error";
    error?: string;
}

export const Edit: PageComponent<Data> = ({addr}, update) => {
    if (!client(addr!).isAuthorized()) {
        showChip("Missing authentication", 4000);
        setTimeout(() => safeRedirect(routes.login, addr!));
        return () => null;
    }

    // Load page data.
    client(addr!).pageRead(
        (data: IPageData) => update({value: data.raw, status: "saved"}),
        () => {
            // TODO handle auth and network errors differently.
            showChip("Unable to verify authentication status", 6000);
            safeRedirect(routes.login, addr!);
        },
    );

    // Save editor contents after a delay.
    // Counter prevents stale requests from updating the status.
    let timeout: any = null;
    let counter = 0;
    const save = (value: string) => {
        update({value, status: "saving"});
        clearTimeout(timeout);
        counter++;
        const selfCounter = counter;
        timeout = setTimeout(() => {
            client(addr!).pageUpdate(
                () => {
                    if (selfCounter !== counter) return;
                    update({value, status: "saved"});
                },
                (m) => {
                    if (selfCounter !== counter) return;
                    update({value, status: "error", error: m});
                },
                value,
            );
        }, saveDelay);
    };

    // Navigate to the document page with keyboard.
    let saving = false;
    keyboard((e) => {
        if (!saving && e.key === "Escape") {
            safeRedirect(routes.document, addr!);
        }
    });

    return (data?: Data) => {
        saving = !!data && data.status === "saving";

        // Response not yet received.
        if (!data) return <Loading />;

        // Change status depending on state.
        let statusContent: any = null;
        if (data.status === "error") {
            statusContent = data.error;
        } else if (data.status === "saving") {
            statusContent = "saving ...";
        } else if (data.status === "saved") {
            statusContent = (
                <span>
                    <Icon name="check" color="yellowgreen" />
                    &nbsp;saved
                </span>
            );
        }

        return (
            <Wrapper>
                <Header>
                    <Done className={{disabled: data.status === "saving"}}>
                        <Anchor id="done" href={path(routes.document, addr!)}>
                            done
                        </Anchor>
                    </Done>
                    <Status className={{error: !!data.error}}>
                        {statusContent}
                    </Status>
                </Header>
                <HeaderPlaceholder />
                <Editor id="page-editor" callback={save} value={data.value} />
            </Wrapper>
        );
    };
};

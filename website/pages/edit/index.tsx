import {client, IPageData} from "../../internal/client";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";
import {Loading} from "../../components/loading";
import {Anchor} from "../../components/anchor";
import {Editor} from "../../components/editor";
import {Icon} from "../../components/icon";
import {keyboard} from "../../internal/keyboard";
import {path, routes, redirect} from "../../routes";

const headerHeight = "2.9rem";
const saveDelay = 1400;

const Wrapper = styled("div")({
    backgroundColor: "#fafafa",
    minHeight: "100%",
});

const Header = styled("header")({
    "-moz-user-select": "-moz-none",
    backgroundColor: "#fff",
    borderBottom: "1px solid #ddd",
    height: headerHeight,
    padding: "0.85rem 1.4rem",
    position: "fixed",
    userSelect: "none",
    width: "100%",
    zIndex: 1,
});

const Done = styled("div")({
    display: "inline-block",
    float: "right",
    fontWeight: "bold",

    $nest: {
        "&.disabled": {
            pointerEvents: "none",
            color: "#ccc",
        },

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

const Status = styled("div")({
    display: "inline-block",
    fontWeight: "bold",

    $nest: {
        "&.error": {
            color: "tomato",
        },
    },
});

interface Data {
    page: IPageData;
    status: "saving" | "saved" | "error";
    error?: string;
}

export const Edit: PageComponent<Data> = ({addr}, update) => {
    // Load page data.
    client.page.fetch(
        (data: IPageData) => update({page: data, status: "saved"}),
        () => redirect(routes.login, addr),
        addr,
    );

    // Save editor contents after a delay.
    // Counter prevents stale requests from updating the status.
    let timeout: any = null;
    let counter = 0;
    const save = (data: Data) => (value) => {
        update({page: data.page, status: "saving"});
        clearTimeout(timeout);
        counter++;
        const selfCounter = counter;
        timeout = setTimeout(() => {
            client.page.edit(
                () => {
                    if (selfCounter !== counter) return;
                    update({page: data.page, status: "saved"});
                },
                (m) => {
                    if (selfCounter !== counter) return;
                    update({page: data.page, status: "error", error: m});
                },
                addr,
                value,
            );
        }, saveDelay);
    };

    // Navigate to the document page with "ctrl+enter".
    keyboard((e) => {
        if (e.ctrl && e.key === "Enter") {
            redirect(routes.document, addr);
        }
    });

    return (data?: Data) => {
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
                        <Anchor href={path(routes.document, addr)}>done</Anchor>
                    </Done>
                    <Status className={{error: !!data.error}}>
                        {statusContent}
                    </Status>
                </Header>
                <Editor
                    id="page-editor"
                    callback={save(data)}
                    value={data.page.raw}
                />
            </Wrapper>
        );
    };
};

import {app} from "../../internal/app";
import {client, IPageData} from "../../internal/client";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";
import {Loading} from "../../components/loading";

const editorID = "targetblank-editor";
const headerHeight = "2.9rem";
const lineHeight = "1.6rem";
const editorPadding = "2.2rem";
const saveDelay = 1400;

const Wrapper = styled("div")({
    backgroundColor: "#fafafa",
    minHeight: "100%",
});

const Header = styled("header")({
    backgroundColor: "#fff",
    borderBottom: "1px solid #ddd",
    height: headerHeight,
    padding: "0.85rem 1.4rem",
    position: "fixed",
    userSelect: "none",
    width: "100%",
    zIndex: 1,
});

const DoneButton = styled("div")({
    cursor: "pointer",
    display: "inline-block",
    float: "right",
    userSelect: "none",
    fontWeight: "bold",
});

const Status = styled("div")({
    display: "inline-block",
    fontWeight: "bold",

    "&.error": {
        color: "tomato",
    },
});

const SavedIcon = styled("i")({
    color: "yellowgreen",
    margin: "0 0.4rem",
    transform: "translate(0, 0.1rem)",
});

const Editor = styled("textarea")({
    lineHeight,
    backgroundColor: "#fafafa",
    border: "none",
    color: "#333",
    fontFamily: "Inconsolata, monospace",
    fontSize: "1.1rem",
    fontWeight: 700,
    marginTop: headerHeight,
    minHeight: "100%",
    outline: "none",
    padding: editorPadding,
    paddingLeft: `calc(2 * ${editorPadding})`,
    paddingRight: `calc(2 * ${editorPadding})`,
    resize: "none",
    whiteSpace: "pre",
    width: "100%",
});

const Lines = styled("div")({
    height: 0,
    opacity: 0.2,
    textAlign: "right",
    transform: `translateY(calc(${headerHeight} + ${editorPadding}))`,
    userSelect: "none",
    width: `calc(1.2 * ${editorPadding})`,
});

const Line = styled("div")({
    lineHeight,
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
        () => app.redirect(`/${addr}/login`),
        addr,
    );

    // Swallow `ctrl+s` to prevent browser popup.
    // https://stackoverflow.com/questions/4446987/overriding-controls-save-functionality-in-browser
    document.addEventListener(
        "keydown",
        (e) => {
            const ctrl = navigator.platform.match("Mac")
                ? e.metaKey
                : e.ctrlKey;
            if (e.key === "s" && ctrl) {
                e.preventDefault();
            }
        },
        false,
    );

    // Save editor contents after a delay.
    // Counter prevents stale requests from updating the status.
    let timeout: any = null;
    let counter = 0;
    const onInput = (data: Data) => (e) => {
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
                e.target.value,
            );
        }, saveDelay);
    };

    // Insert spaces when tab is pressed.
    const onKeydown = (e) => {
        if (e.key === "Tab") {
            e.preventDefault();
            const {target} = e;
            const pos = target.selectionStart;
            const before = target.value.substring(0, target.selectionStart);
            const after = target.value.substring(target.selectionEnd);
            target.value = `${before}    ${after}`;
            target.selectionEnd = pos + 4;
        }
    };

    return (data?: Data) => {
        // Response not yet received.
        if (!data) return <Loading />;

        // Update editor height to match content.
        const editor = document.getElementById(editorID);
        if (editor) {
            editor.style.height = "0";
            editor.style.opacity = "1";
            editor.style.height = `${editor.scrollHeight + 20}px`;
        }

        // Create line numbers.
        let lines: number[] = [];
        if (editor) {
            const count = (editor as any).value.split("\n").length;
            lines = Array(count);
            for (let i = 0; i < count; i++) {
                lines[i] = i + 1;
            }
        }

        // Change status depending on state.
        let statusContent: any = null;
        if (data.status === "error") {
            statusContent = data.error;
        } else if (data.status === "saving") {
            statusContent = "saving ...";
        } else if (data.status === "saved") {
            statusContent = (
                <span>
                    <SavedIcon className="far fa-lg fa-check-circle" />
                    saved
                </span>
            );
        }

        return (
            <Wrapper>
                <Header>
                    <DoneButton onclick={() => app.redirect(`/${addr}`)}>
                        done
                    </DoneButton>
                    <Status className={{error: !!data.error}}>
                        {statusContent}
                    </Status>
                </Header>
                <Lines>{...lines.map((n) => <Line>{n}</Line>)}</Lines>
                <Editor
                    id={editorID}
                    style="opacity: 0;"
                    value={data.page.raw}
                    oninput={onInput(data)}
                    onkeydown={onKeydown}
                    spellcheck={false}
                />
            </Wrapper>
        );
    };
};

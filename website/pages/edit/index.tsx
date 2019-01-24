import {app} from "../../internal/app";
import {client, IPageData} from "../../internal/client";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";

const editorID = "targetblank-editor";
const headerHeight = "2.9rem";
const lineHeight = "1.6rem";
const editorPadding = "2.2rem";
const saveDelay = 1400;

const Wrapper = styled("div")({
    minHeight: "100%",
    backgroundColor: "#fafafa",
});

const Header = styled("header")({
    backgroundColor: "#fff",
    borderBottom: "1px solid #ddd",
    height: headerHeight,
    padding: "0.85rem 1.4rem",
    position: "fixed",
    userSelect: "none",
    width: "100%",
});

const BackButton = styled("div")({
    cursor: "pointer",
    display: "inline-block",
    userSelect: "none",
    fontWeight: "bold",
});

const Status = styled("div")({
    display: "inline-block",
    float: "right",
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
    fontFamily: "Inconsolata, monospace",
    fontSize: "1.2rem",
    marginTop: headerHeight,
    outline: "none",
    padding: editorPadding,
    paddingLeft: `calc(2 * ${editorPadding})`,
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
    client.page.fetch(
        (data: IPageData) => update({page: data, status: "saved"}),
        () => app.redirect(`/${addr}/login`),
        addr,
    );

    return (data?: Data) => {
        // Response not yet received.
        if (!data) {
            // TODO
            return "loading";
        }

        // Save editor contents after a delay.
        let timeout: any = null;
        const onInput = (e) => {
            update({page: data.page, status: "saving"});
            clearTimeout(timeout);
            timeout = setTimeout(() => {
                client.page.edit(
                    () => update({page: data.page, status: "saved"}),
                    (m) => update({page: data.page, status: "error", error: m}),
                    addr,
                    e.target.value,
                );
            }, saveDelay);
        };

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
                    <BackButton onclick={() => app.redirect(`/${addr}`)}>
                        back
                    </BackButton>
                    <Status className={{error: !!data.error}}>
                        {statusContent}
                    </Status>
                </Header>
                <Lines>{...lines.map((n) => <Line>{n}</Line>)}</Lines>
                <Editor
                    id={editorID}
                    style="opacity: 0;"
                    value={data.page.raw}
                    oninput={onInput}
                    onkeydown={onKeydown}
                    spellcheck={false}
                />
            </Wrapper>
        );
    };
};

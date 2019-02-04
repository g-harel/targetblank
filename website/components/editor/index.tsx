import {Component} from "../../internal/types";
import {styled} from "../../internal/styled";

const headerHeight = "2.9rem";
const lineHeight = "1.6rem";
const editorPadding = "1.8rem";

const Wrapper = styled("div")({});

const Textarea = styled("textarea")({
    lineHeight,
    backgroundColor: "#fafafa",
    border: "none",
    color: "#333",
    fontFamily:
        "SFMono-Regular, Consolas, Liberation Mono, Menlo, Courier, monospace",
    marginTop: headerHeight,
    minHeight: "100%",
    outline: "none",
    paddingTop: editorPadding,
    paddingBottom: 0,
    paddingLeft: `calc(2.4 * ${editorPadding})`,
    paddingRight: `calc(2.4 * ${editorPadding})`,
    resize: "none",
    whiteSpace: "pre",
    width: "100%",
});

const Lines = styled("div")({
    height: 0,
    color: "#ccc",
    textAlign: "right",
    transform: `translateY(calc(${headerHeight} + ${editorPadding} + 1px))`,
    userSelect: "none",
    width: `calc(1.6 * ${editorPadding} + 1rem)`,
});

const Line = styled("div")({
    lineHeight,
    backgroundColor: "#fafafa",
    paddingRight: "1rem",
});

export interface Props {
    id: string;
    callback: (input: string) => any;
    value: string;
}

export const Editor: Component<Props> = (props) => () => {
    // Set focus to start of textarea.
    setTimeout(() =>
        requestAnimationFrame(() => {
            const input = document.getElementById(props.id);
            if (input && document.activeElement !== input) {
                input.focus();
                (input as any).setSelectionRange(0, 0);
            }
        }),
    );

    const onKeydown = (e) => {
        // Swallow `ctrl+s` to prevent browser popup.
        const ctrl = navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey;
        if (ctrl && e.key === "s") {
            e.preventDefault();
        }

        // Insert spaces when tab is pressed.
        if (e.key === "Tab") {
            e.preventDefault();
            const {target} = e;
            const pos = target.selectionStart;
            const before = target.value.substring(0, target.selectionStart);
            const after = target.value.substring(target.selectionEnd);
            target.value = `${before}    ${after}`;
            target.selectionEnd = pos + 4;
            props.callback(target.value);
        }
    };

    const editor = document.getElementById(props.id);

    // Update editor height to match content.
    if (editor) {
        editor.style.height = "0";
        editor.style.opacity = "1";
        editor.style.height = `${editor.scrollHeight + 20}px`;
        editor.style.marginBottom = "-20px";
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

    return (
        <Wrapper>
            <Lines>{...lines.map((n) => <Line>{n}</Line>)}</Lines>
            <Textarea
                id={props.id}
                style="opacity: 0;"
                value={props.value}
                oninput={(e) => props.callback(e.target.value)}
                onkeydown={onKeydown}
                spellcheck={false}
            />
        </Wrapper>
    );
};

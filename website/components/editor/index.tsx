import {Component} from "../../internal/types";
import {styled, colors, fonts} from "../../internal/style";
import {FileEditor} from "./file";

const headerHeight = "2.9rem";
const lineHeight = "1.6rem";
const editorPadding = "1.8rem";

const Wrapper = styled("div")({});

const StyledTextarea = styled("textarea")({
    lineHeight,
    backgroundColor: colors.backgroundSecondary,
    border: "none",
    color: colors.textPrimary,
    fontFamily: fonts.monospace,
    marginTop: headerHeight,
    minHeight: "100%",
    outline: "none",
    overflowY: "hidden",
    paddingLeft: `calc(2.4 * ${editorPadding})`,
    paddingTop: editorPadding,
    resize: "none",
    whiteSpace: "pre",
    width: "100%",
});

const Lines = styled("div")({
    "-moz-user-select": "none",
    height: 0,
    color: colors.textSecondarySmall,
    textAlign: "right",
    transform: `translateY(calc(${headerHeight} + ${editorPadding}))`,
    userSelect: "none",
    width: `calc(1.6 * ${editorPadding} + 1rem)`,
});

const Line = styled("div")({
    lineHeight,
    backgroundColor: colors.backgroundSecondary,
    paddingRight: "1rem",
});

export interface Props {
    id: string;
    callback: (input: string) => any;
    value: string;
}

// TODO Add shortcuts to docs (readme + new page)
export const Editor: Component<Props> = (props) => () => {
    const onKeydown = (e: KeyboardEvent) => {
        updateCursorPosition(e);

        // Swallow `ctrl+s` to prevent browser popup.
        const ctrl = navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey;
        if (ctrl && e.key === "s") {
            e.preventDefault();
        }

        // Standardized helper to modify the contents in response to a command.
        const editFile = (op: keyof FileEditor) => {
            e.preventDefault();

            const target = (e.target as any) as HTMLTextAreaElement;
            const initialValue = target.value;
            const fileEditor = new FileEditor(
                target.value,
                target.selectionStart,
                target.selectionEnd,
            );

            fileEditor[op]();

            target.value = fileEditor.getContent();
            target.selectionStart = fileEditor.getSelectionStart();
            target.selectionEnd = fileEditor.getSelectionEnd();

            if (initialValue !== target.value) {
                props.callback(target.value);
            }
        };

        // Modify indentation when tab is pressed.
        if (e.key === "Tab") {
            if (e.shiftKey) {
                editFile("unindent");
            } else {
                editFile("indent");
            }
        }

        // Move lines using alt + arrows.
        if (e.key === "ArrowUp" && e.altKey) {
            editFile("moveUp");
        }
        if (e.key === "ArrowDown" && e.altKey) {
            editFile("moveDown");
        }
    };

    const updateCursorPosition = (e: any) => {
        (window as any).editor = (window as any).editor || {};
        (window as any).editor[props.id] = e.target.selectionStart;
    };

    const onInput = (e: any) => {
        props.callback(e.target.value);
    };

    setTimeout(() =>
        requestAnimationFrame(() => {
            const editor = document.getElementById(props.id);

            // Set focus to last known position.
            if (editor && document.activeElement !== editor) {
                editor.focus();
                const position = ((window as any).editor || {})[props.id];
                (editor as any).setSelectionRange(position, position);
            }

            // Update editor height to match content.
            if (editor) {
                editor.style.opacity = "1";
                editor.style.height = `${editor.scrollHeight}px`;
            }
        }),
    );

    // Create line numbers.
    let lines: number[] = [];
    const count = props.value.split("\n").length;
    lines = Array(count);
    for (let i = 0; i < count; i++) {
        lines[i] = i + 1;
    }

    return (
        <Wrapper>
            <Lines>{...lines.map((n) => <Line>{n}</Line>)}</Lines>
            <StyledTextarea
                id={props.id}
                style="opacity: 0;"
                value={props.value}
                oninput={onInput}
                onkeydown={onKeydown}
                onclick={updateCursorPosition}
                spellcheck={false}
            />
        </Wrapper>
    );
};

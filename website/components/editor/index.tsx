import {Component} from "../../internal/types";
import {styled, colors, fonts} from "../../internal/style";
import {FileEditor} from "../../internal/editor/file";

// Values used to compute the approximate size of the the textbox without hacks.
// https://stackoverflow.com/q/2803880
const lineHeight = 1.6;
const charWidthRatio = 0.66;

const Wrapper = styled("div")({
    display: "flex",
    flexDirection: "row",
    padding: "1.8rem 0 0 1.8rem",
});

const LineNumbers = styled("div")({
    display: "flex",
    flexDirection: "column",
    paddingRight: "1rem",
});

const LineNumber = styled("div")({
    "-moz-user-select": "none",
    backgroundColor: colors.backgroundSecondary,
    color: colors.textSecondarySmall,
    lineHeight: `${lineHeight}rem`,
    textAlign: "right",
    userSelect: "none",
});

// Wrapper to prevent the scrollbar from moving the content on chrome.
const ScrollWrapper = styled("div")({
    flexGrow: 1,
    overflowX: "auto",
});

const StyledTextarea = styled("textarea")({
    backgroundColor: colors.backgroundSecondary,
    border: "none",
    color: colors.textPrimary,
    fontFamily: fonts.monospace,
    lineHeight: `${lineHeight}rem`,
    outline: "none",
    overflow: "hidden",
    resize: "none",
    whiteSpace: "pre",
});

export interface Props {
    id: string;
    callback: (input: string) => any;
    value: string;
}

// Editor component holds no state and expects its parent
// to update the value as provided to the callback. This
// ensures the editor is only updated once on each change
// since the parent is expected to already get re-rendered.
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
            } else if (!e.altKey) {
                editFile("indent");
            }
        }

        // Indent lines using ctrl + brackets.
        if (ctrl) {
            if (e.key === "[") editFile("unindent");
            if (e.key === "]") editFile("indent");
            if (e.key === "/") editFile("toggleComment");
        }

        // Move lines using alt + arrows.
        if (e.altKey) {
            if (e.key === "ArrowUp") editFile("moveUp");
            if (e.key === "ArrowDown") editFile("moveDown");
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
                editor.focus({preventScroll: true});
                const position = ((window as any).editor || {})[props.id];
                (editor as any).setSelectionRange(position, position);
            }
        }),
    );

    const lines = props.value.split("\n");

    // Find longest line.
    let longest = 0;
    for (const line of lines) {
        longest = Math.max(longest, line.length);
    }

    // Create line numbers.
    let lineNumbers: number[] = [];
    const count = lines.length;
    lineNumbers = Array(count);
    for (let i = 0; i < count; i++) {
        lineNumbers[i] = i + 1;
    }

    const style = `
        width: ${charWidthRatio * longest + 3}em;
        height: ${lineHeight * lines.length}em;
    `;

    return (
        <Wrapper>
            <LineNumbers>
                {...lineNumbers.map((n) => <LineNumber>{n}</LineNumber>)}
            </LineNumbers>
            <ScrollWrapper>
                <StyledTextarea
                    id={props.id}
                    style={style}
                    value={props.value}
                    oninput={onInput}
                    onkeydown={onKeydown}
                    onclick={updateCursorPosition}
                    spellcheck={false}
                />
            </ScrollWrapper>
        </Wrapper>
    );
};

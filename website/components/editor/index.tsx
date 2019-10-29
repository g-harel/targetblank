import {Component} from "../../internal/types";
import {styled} from "../../internal/style";
import {color, font} from "../../internal/style/theme";
import {Command, command} from "../../internal/editor";

// Values used to compute the approximate size of the the textbox without hacks.
// https://stackoverflow.com/q/2803880
const lineHeight = 1.6;
const charWidthRatio = 0.66;

const Wrapper = styled("div")({
    display: "flex",
    flexDirection: "row",
    flexGrow: 1,
    padding: "1.8rem 0 0 1.8rem",
});

const LineNumbers = styled("div")({
    display: "flex",
    flexDirection: "column",
    paddingRight: "1rem",
});

const LineNumber = styled("div")({
    "-moz-user-select": "none",
    backgroundColor: color.backgroundSecondary,
    color: color.textSecondarySmall,
    lineHeight: `${lineHeight}rem`,
    minWidth: "2em",
    textAlign: "right",
    userSelect: "none",
});

const Whitespace = styled("div")({
    display: "flex",
    flexDirection: "column",
    height: 0,
    width: 0,
    zIndex: 2,
});

const WhitespaceLine = styled("div")({
    "-moz-user-select": "none",
    color: color.textPrimary,
    fontFamily: font.monospace,
    lineHeight: `${lineHeight}rem`,
    // Use opacity and primary text color for whitespace overlay to be invisible
    // when document content contains whitespace character. A single-phase
    // character replacement pass would remove the need for this workaround.
    opacity: 0.2,
    pointerEvents: "none",
    userSelect: "none",
    whiteSpace: "pre",
});

// Wrapper to prevent the scrollbar from moving the content on chrome.
const ScrollWrapper = styled("div")({
    display: "flex",
    flexDirection: "column",
    flexGrow: 1,
    overflowX: "auto",
});

const StyledTextarea = styled("textarea")({
    "-moz-tab-size": 1,
    backgroundColor: color.backgroundSecondary,
    border: "none",
    color: color.textPrimary,
    flexGrow: 1,
    fontFamily: font.monospace,
    lineHeight: `${lineHeight}rem`,
    minWidth: "100%",
    outline: "none",
    overflow: "hidden",
    resize: "none",
    // Characters of width more than one break white space rendering.
    tabSize: 1,
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
        const editFile = (command: Command) => {
            e.preventDefault();

            const target = (e.target as any) as HTMLTextAreaElement;
            const initialValue = target.value;

            const out = command(target);
            target.value = out.value;
            target.selectionStart = out.selectionStart;
            target.selectionEnd = out.selectionEnd;

            if (initialValue !== target.value) {
                props.callback(target.value);
            }
        };

        // Modify indentation when tab is pressed.
        if (e.key === "Tab") {
            if (e.shiftKey) {
                editFile(command.unindent);
            } else if (!e.altKey) {
                editFile(command.indent);
            }
        }

        // Indent lines using ctrl + brackets.
        if (ctrl) {
            if (e.key === "[") editFile(command.unindent);
            if (e.key === "]") editFile(command.indent);
            if (e.key === "/") editFile(command.toggleComment);
        }

        if (e.key === "Enter") {
            if (!ctrl && !e.altKey && !e.shiftKey) {
                editFile(command.newline);
            }
        }

        // Move lines using alt + arrows.
        if (e.altKey) {
            if (e.key === "ArrowUp") editFile(command.moveUp);
            if (e.key === "ArrowDown") editFile(command.moveDown);
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
                <Whitespace>
                    {...lines.map((l) => (
                        <WhitespaceLine>
                            {l.replace(/ /g, "·").replace(/[^·]/g, " ")}&nbsp;
                        </WhitespaceLine>
                    ))}
                </Whitespace>
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

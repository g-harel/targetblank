import {Component} from "../../internal/types";
import {styled, colors, fonts} from "../../internal/style";

const headerHeight = "2.9rem";
const lineHeight = "1.6rem";
const editorPadding = "1.8rem";
const tab = "    ";

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

export const Editor: Component<Props> = (props) => () => {
    const onKeydown = (e: KeyboardEvent) => {
        updateCursorPosition(e);

        // Swallow `ctrl+s` to prevent browser popup.
        const ctrl = navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey;
        if (ctrl && e.key === "s") {
            e.preventDefault();
        }

        // Insert spaces when tab is pressed.
        if (e.key === "Tab") {
            e.preventDefault();
            const target = (e.target as any) as HTMLTextAreaElement;
            const pos = target.selectionStart;
            const before = target.value.substring(0, target.selectionStart);
            const after = target.value.substring(target.selectionEnd);

            if (e.shiftKey) {
                const beforeLines = before.split("\n");
                if (beforeLines.length > 0) {
                    const lastLineIndex = beforeLines.length - 1;
                    const lastLine = beforeLines[lastLineIndex];
                    if (lastLine.startsWith(tab)) {
                        beforeLines[lastLineIndex] = lastLine.substr(
                            tab.length,
                        );
                        target.value = `${beforeLines.join("\n")}${after}`;
                        target.selectionStart = pos - 4;
                        target.selectionEnd = pos - 4;
                    }
                }
            } else {
                target.value = `${before}${tab}${after}`;
                target.selectionEnd = pos + 4;
            }

            props.callback(target.value);
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

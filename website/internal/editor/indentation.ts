import {Command} from ".";
import {
    init,
    lineByPos,
    lineIsEmpty,
    isMultilineSelection,
    leadingSpace,
    render,
    posByLine,
    indexInLine,
} from "./util";

export const INDENT = "    ";
export const INDENT_LENGTH = INDENT.length;

// Increase indentation level of all selected lines.
// Cursor positions are updated.
export const indent: Command = (state) => {
    const s = init(state);

    const selectionStartLine = lineByPos(s, s.selectionStart);
    const selectionEndLine = lineByPos(s, s.selectionEnd);

    // Modify line contents.
    let firstLineAdded = 0;
    let totalAdded = 0;
    for (let i = selectionStartLine; i <= selectionEndLine; i++) {
        if (lineIsEmpty(s, i) && isMultilineSelection(s)) continue;
        const add = INDENT_LENGTH - (leadingSpace(s, i) % INDENT_LENGTH);
        s.lines[i] = INDENT.slice(0, add) + s.lines[i];
        if (i === selectionStartLine) firstLineAdded = add;
        totalAdded += add;
    }

    // Modify selection indecies.
    s.selectionStart += firstLineAdded;
    s.selectionEnd += totalAdded;

    return render(s);
};

// Decrease indentation level of all selected lines.
// Cursor positions are updated, but will not move across lines.
export const unindent: Command = (state) => {
    const s = init(state);

    const selectionStartLine = lineByPos(s, s.selectionStart);
    const selectionEndLine = lineByPos(s, s.selectionEnd);

    // Modify line contents (if they start with an indent).
    let firstLineRemoved = 0;
    let totalRemoved = 0;
    for (let i = selectionStartLine; i <= selectionEndLine; i++) {
        if (lineIsEmpty(s, i) && isMultilineSelection(s)) continue;
        const whitespace = leadingSpace(s, i);
        if (whitespace === 0) continue;
        const remove = whitespace % INDENT_LENGTH || INDENT_LENGTH;
        s.lines[i] = s.lines[i].substr(remove);
        if (i === selectionStartLine) firstLineRemoved = remove;
        totalRemoved += remove;
    }

    // Modify selection indecies.
    s.selectionStart -= firstLineRemoved;
    s.selectionEnd -= totalRemoved;

    // Keep cursor on same line it was before unindent.
    if (selectionStartLine !== lineByPos(s, s.selectionStart)) {
        s.selectionStart = posByLine(s, selectionStartLine);
    }
    if (selectionEndLine !== lineByPos(s, s.selectionEnd)) {
        s.selectionEnd = posByLine(s, selectionEndLine);
    }

    return render(s);
};

export const newline: Command = (state) => {
    const s = init(state);

    const selectionLine = lineByPos(s, s.selectionEnd);
    const cursorPositionInLine = indexInLine(s, s.selectionEnd);

    const whitespace = leadingSpace(s, selectionLine);
    const before = s.lines[selectionLine].slice(0, cursorPositionInLine);
    const after = s.lines[selectionLine].slice(cursorPositionInLine);
    const lineIndex = selectionLine + 1;

    s.lines[selectionLine] = before;
    s.lines.splice(lineIndex, 0, before.slice(0, whitespace) + after);

    const newSelection =
        posByLine(s, lineIndex) + Math.min(whitespace, before.length);
    s.selectionStart = newSelection;
    s.selectionEnd = newSelection;

    return render(s);
};

import {Command} from ".";
import {INDENT_LENGTH} from "./indentation";
import {
    init,
    lineByPos,
    lineIsEmpty,
    isMultilineSelection,
    leadingSpace,
    render,
    indexInLine,
} from "./util";

export const COMMENT = "#";
export const COMMENT_LENGTH = COMMENT.length;
export const COMMENT_REPLACE_PATTERN = /^(\s*)# ?/g;

export const toggleComment: Command = (state) => {
    const s = init(state);

    const selectionStartLine = lineByPos(s, s.selectionStart);
    const selectionEndLine = lineByPos(s, s.selectionEnd);

    // Decide whether to comment or un-comment selection. Lines will be
    // commented out if any of the selected lines is not currently commented.
    let comment = false;
    for (let i = selectionStartLine; i <= selectionEndLine; i++) {
        if (lineIsEmpty(s, i) && isMultilineSelection(s)) continue;
        comment = comment || !s.lines[i].trim().startsWith(COMMENT);
    }

    return comment ? commentOn(state) : commentOff(state);
};

const commentOn: Command = (state) => {
    const s = init(state);

    const selectionStartLine = lineByPos(s, s.selectionStart);
    const selectionEndLine = lineByPos(s, s.selectionEnd);
    const startIndexInLine = indexInLine(s, s.selectionStart);
    const endIndexInLine = indexInLine(s, s.selectionEnd);

    // Find highest possible level that can be commented.
    let level = Infinity;
    for (let i = selectionStartLine; i <= selectionEndLine; i++) {
        if (lineIsEmpty(s, i) && isMultilineSelection(s)) {
            continue;
        }
        level = Math.min(level, leadingSpace(s, i));
    }

    // Only comment at valid indentation levels.
    level = level - (level % INDENT_LENGTH);
    let totalAdded = 0;
    for (let i = selectionStartLine; i <= selectionEndLine; i++) {
        if (lineIsEmpty(s, i) && isMultilineSelection(s)) {
            continue;
        }
        const line = s.lines[i];
        s.lines[i] = `${line.slice(0, level)}${COMMENT} ${line.slice(level)}`;
        totalAdded += COMMENT_LENGTH + 1;
    }

    // Modify selection indecies.
    if (startIndexInLine >= level) {
        s.selectionStart += COMMENT_LENGTH + 1;
    }
    if (endIndexInLine >= level) {
        s.selectionEnd += totalAdded;
    } else {
        s.selectionEnd += totalAdded - COMMENT_LENGTH - 1;
    }

    return render(s);
};

const commentOff: Command = (state) => {
    const s = init(state);

    const selectionStartLine = lineByPos(s, s.selectionStart);
    const selectionEndLine = lineByPos(s, s.selectionEnd);
    const startIndexInLine = indexInLine(s, s.selectionStart);
    const endIndexInLine = indexInLine(s, s.selectionEnd);

    let firstLineRemoved = 0;
    let firstLineRemovedIndex = 0;
    let lastLineRemoved = 0;
    let lastLineRemovedIndex = 0;
    let totalRemoved = 0;
    for (let i = selectionStartLine; i <= selectionEndLine; i++) {
        if (lineIsEmpty(s, i) && isMultilineSelection(s)) {
            continue;
        }
        if (i === selectionStartLine) {
            firstLineRemovedIndex = leadingSpace(s, i);
        }
        if (i === selectionEndLine - 1) {
            lastLineRemovedIndex = leadingSpace(s, i);
        }
        const line = s.lines[i];
        s.lines[i] = line.replace(COMMENT_REPLACE_PATTERN, "$1");
        const removed = line.length - s.lines[i].length;
        totalRemoved += removed;
        if (i === selectionStartLine) firstLineRemoved = removed;
        if (i === selectionEndLine - 1) lastLineRemoved = removed;
    }

    // Modify selection indecies.
    if (startIndexInLine >= firstLineRemovedIndex) {
        s.selectionStart -= firstLineRemoved;
    }
    if (endIndexInLine >= lastLineRemovedIndex) {
        s.selectionEnd -= totalRemoved;
    } else {
        s.selectionEnd -= totalRemoved - lastLineRemoved;
    }

    return render(s);
};

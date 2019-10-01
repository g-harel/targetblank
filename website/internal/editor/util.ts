import {EditorState, InternalEditorState} from ".";

// Initializes the internal editor state. Returned value is safe to modify
// without affecting original editor state.
export const init = (s: EditorState): InternalEditorState => {
    // Selection indecies are checked on creation.
    // Methods that modify the file will keep the selection valid.
    if (s.selectionStart > s.value.length || s.selectionStart < 0) {
        throw new Error("FileEditor: selection start index is out of range.");
    }
    if (s.selectionEnd > s.value.length || s.selectionEnd < 0) {
        throw new Error("FileEditor: selection end index is out of range.");
    }
    if (s.selectionStart > s.selectionEnd) {
        throw new Error(
            "FileEditor: selection start position is higher than end",
        );
    }
    const lines = s.value.split("\n");
    return {
        lines,
        selectionStart: s.selectionStart,
        selectionEnd: s.selectionEnd,
    };
};

// Converts the internal state back to an editor state.
export const render = (s: InternalEditorState): EditorState => {
    return {
        value: s.lines.join("\n"),
        selectionStart: s.selectionStart,
        selectionEnd: s.selectionEnd,
    };
};

// Computes the line number at which the index is found.
// Line breaks are counted as a single character.
export const lineByPos = (s: InternalEditorState, pos: number): number => {
    let charCount = 0;
    let lineCount = 0;
    for (const line of s.lines) {
        charCount += line.length + 1;
        if (charCount > pos) {
            break;
        }
        lineCount++;
    }
    return lineCount;
};

// Computes the index of the first character on the given line.
// Line breaks are counted as a single character.
export const posByLine = (s: InternalEditorState, line: number): number => {
    let charCount = 0;
    for (let i = 0; i < line; i++) {
        charCount += s.lines[i].length + 1;
    }
    return charCount;
};

// Decides whether or not a line contains content.
export const lineIsEmpty = (s: InternalEditorState, line: number): boolean => {
    return s.lines[line].trim() === "";
};

// Computes whether the selection is multiline or not.
export const isMultilineSelection = (s: InternalEditorState): boolean => {
    return lineByPos(s, s.selectionStart) !== lineByPos(s, s.selectionEnd);
};

// Calculates the number of spaces at the start of the line.
export const leadingSpace = (s: InternalEditorState, line: number): number => {
    const contents = s.lines[line];
    let count = 0;
    while (contents[count] === " ") count++;
    return count;
};

// Computes the cursor position of the given position in its line.
export const indexInLine = (s: InternalEditorState, pos: number): number => {
    return pos - posByLine(s, lineByPos(s, pos));
};

import {Command} from ".";
import {init, lineByPos, render} from "./util";

// Moves the currently selected lines up by one.
export const moveUp: Command = (state) => {
    const s = init(state);

    const selectionStartLine = lineByPos(s, s.selectionStart);
    const selectionEndLine = lineByPos(s, s.selectionEnd);
    if (selectionStartLine === 0) return render(s);

    // Swap lines in place till the entire selection is moved.
    for (let i = selectionStartLine; i <= selectionEndLine; i++) {
        const temp = s.lines[i];
        s.lines[i] = s.lines[i - 1];
        s.lines[i - 1] = temp;
    }

    // Modify selection indecies.
    const movedChars = s.lines[selectionEndLine].length + 1;
    s.selectionStart -= movedChars;
    s.selectionEnd -= movedChars;

    return render(s);
};

// Moves the currently selected lines down by one.
export const moveDown: Command = (state) => {
    const s = init(state);

    const selectionStartLine = lineByPos(s, s.selectionStart);
    const selectionEndLine = lineByPos(s, s.selectionEnd);
    if (selectionEndLine >= s.lines.length) return render(s);

    // Swap lines in place till the entire selection is moved.
    for (let i = selectionEndLine; i >= selectionStartLine; i--) {
        const temp = s.lines[i];
        s.lines[i] = s.lines[i + 1];
        s.lines[i + 1] = temp;
    }

    // Modify selection indecies.
    const movedChars = s.lines[selectionStartLine].length + 1;
    s.selectionStart += movedChars;
    s.selectionEnd += movedChars;

    return render(s);
};

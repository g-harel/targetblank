import {EditorState} from ".";

export const genRandomState = (size: number): EditorState => {
    const lineCount = Math.ceil(size * Math.random());
    const lines = [];
    for (let j = 0; j < lineCount; j++) {
        const whitespaceCount = Math.floor(size * Math.random());
        const contentCount = Math.ceil(size * Math.random());
        lines.push(" ".repeat(whitespaceCount) + "a".repeat(contentCount));
    }
    const value = lines.join("\n");
    const a = Math.floor(value.length * Math.random());
    const b = Math.floor(value.length * Math.random());

    return {
        value,
        selectionStart: Math.min(a, b),
        selectionEnd: Math.max(a, b),
    };
};

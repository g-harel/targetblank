import {EditorState, Command} from ".";

const toString = (state: EditorState): string => {
    return `# selectionStart: ${state.selectionStart}
# selectionEnd: ${state.selectionEnd}
# =================
${state.value}`;
};

// Assert that two states are equal.
export const expectEqual = (actual: EditorState, expected: EditorState) => {
    expect(toString(actual)).toBe(toString(expected));
};

// Stateful assertion helper to make command tests easier to read.
export const expectCommand = (c: Command) => {
    const initial: Partial<EditorState> = {};

    let produced: null | EditorState = null;
    const produceState = (): EditorState => {
        if (produced !== null) return produced;
        if (initial.value === undefined) {
            throw "unset initial value";
        }
        if (initial.selectionStart === undefined) {
            throw "unset initial selectionStart";
        }
        if (initial.selectionEnd === undefined) {
            throw "unset initial selectionEnd";
        }
        produced = c(initial as EditorState);
        return produced;
    };

    const withValue = (value: string) => {
        initial.value = value;
        return {withCursor, withSelection};
    };

    const withCursor = (cursor: number) => {
        initial.selectionStart = cursor;
        initial.selectionEnd = cursor;
        return {toProduceValue, toProduceUnchanged, toProduceCursor};
    };

    const withSelection = (start: number, end: number) => {
        initial.selectionStart = start;
        initial.selectionEnd = end;
        return {toProduceValue, toProduceUnchanged, toProduceSelection};
    };

    const toProduceValue = (value: string) => {
        expect(produceState().value).toBe(value);
    };

    const toProduceCursor = (cursor: number) => {
        expect(produceState().selectionStart).toBe(cursor);
        expect(produceState().selectionEnd).toBe(cursor);
    };

    const toProduceSelection = (start: number, end: number) => {
        expect(produceState().selectionStart).toBe(start);
        expect(produceState().selectionEnd).toBe(end);
    };

    const toProduceUnchanged = () => {
        expect(produceState().value).toBe(initial.value);
        expect(produceState().selectionStart).toBe(initial.selectionStart);
        expect(produceState().selectionEnd).toBe(initial.selectionEnd);
    };

    return {withValue};
};

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

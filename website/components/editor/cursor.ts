let editorCursorPositions: Record<string, number | undefined> = {};

export const updateCursorPosition = (editorId: string) => (e: any) => {
    editorCursorPositions[editorId] = e.target.selectionStart;
};

export const readCursorPosition = (editorId: string) => {
    return editorCursorPositions[editorId];
};

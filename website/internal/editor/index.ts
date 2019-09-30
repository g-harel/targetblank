export interface EditorState {
    value: string;
    selectionStart: number;
    selectionEnd: number;
}

// Transformed state containing computed fields.
export interface InternalEditorState {
    lines: string[];
    selectionStart: number;
    selectionEnd: number;
}

export interface Command {
    (initial: EditorState): EditorState;
}

import {toggleComment} from "./comment";
import {indent, unindent, newline} from "./indentation";
import {moveUp, moveDown} from "./movement";

export const command = {
    toggleComment,
    indent,
    unindent,
    newline,
    moveUp,
    moveDown,
};

// Type check.
const _: Record<string, Command> = command;
if (0) console.log(_);

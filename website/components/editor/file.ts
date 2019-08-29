const INDENT = "    ";
const INDENT_LENGTH = INDENT.length;

// Helper class used with methods that edit blocks of text.
// Instances of the class are stateful and will maintain file contents and
// cursor positions.
export class FileEditor {
    private lines: string[];
    private selectionStart: number;
    private selectionEnd: number;

    constructor(content: string, selectionStart: number, selectionEnd: number) {
        // Selection indecies are checked on creation.
        // Methods that modify the file will keep the selection valid.
        if (selectionStart > content.length || selectionStart < 0) {
            throw new Error(
                "FileEditor: selection start index is out of range.",
            );
        }
        if (selectionEnd > content.length || selectionEnd < 0) {
            throw new Error("FileEditor: selection end index is out of range.");
        }
        if (selectionStart > selectionEnd) {
            throw new Error(
                "FileEditor: selection start position is higher than end",
            );
        }

        this.lines = content.split("\n");
        this.selectionStart = selectionStart;
        this.selectionEnd = selectionEnd;
    }

    // Computes the line number at which the index is found.
    // Line breaks are counted as a single character.
    private lineByPos(pos: number): number {
        let charCount = 0;
        let lineCount = 0;
        for (const line of this.lines) {
            charCount += line.length + 1;
            if (charCount > pos) {
                break;
            }
            lineCount++;
        }
        return lineCount;
    }

    // Computes the index of the first character on the given line.
    // Line breaks are counted as a single character.
    private posByLine(line: number): number {
        let charCount = 0;
        for (let i = 0; i < line; i++) {
            charCount += this.lines[i].length + 1;
        }
        return charCount;
    }

    // Increase indentation level of all selected lines.
    // Cursor positions are updated.
    public indent() {
        const selectionStartLine = this.lineByPos(this.selectionStart);
        const selectionEndLine = this.lineByPos(this.selectionEnd);

        // Modify line contents.
        for (let i = selectionStartLine; i <= selectionEndLine; i++) {
            this.lines[i] = INDENT + this.lines[i];
        }

        // Modify selection indecies.
        this.selectionStart += INDENT_LENGTH;
        const editedLines = selectionEndLine - selectionStartLine + 1;
        this.selectionEnd += INDENT_LENGTH * editedLines;
    }

    // Decrease indentation level of all selected lines.
    // Cursor positions are updated, but will not move across lines.
    public unindent() {
        const selectionStartLine = this.lineByPos(this.selectionStart);
        const selectionEndLine = this.lineByPos(this.selectionEnd);

        // Modify line contents (if they start with an indent).
        let editedLines = 0;
        let firstLineEdited = false;
        for (let i = selectionStartLine; i <= selectionEndLine; i++) {
            if (this.lines[i].startsWith(INDENT)) {
                if (i === selectionStartLine) firstLineEdited = true;
                this.lines[i] = this.lines[i].substr(INDENT_LENGTH);
                editedLines++;
            }
        }

        // Modify selection indecies.
        if (firstLineEdited) {
            this.selectionStart -= INDENT_LENGTH;
        }
        this.selectionEnd -= INDENT_LENGTH * editedLines;

        // Keep cursor on same line it was before unindent.
        if (selectionStartLine !== this.lineByPos(this.selectionStart)) {
            this.selectionStart = this.posByLine(selectionStartLine);
        }
        if (selectionEndLine !== this.lineByPos(this.selectionEnd)) {
            this.selectionEnd = this.posByLine(selectionEndLine);
        }
    }

    // Moves the currently selected lines up by one.
    public moveUp() {
        const selectionStartLine = this.lineByPos(this.selectionStart);
        const selectionEndLine = this.lineByPos(this.selectionEnd);
        if (selectionStartLine === 0) return;

        // Swap lines in place till the entire selection is moved.
        for (let i = selectionStartLine; i <= selectionEndLine; i++) {
            const temp = this.lines[i];
            this.lines[i] = this.lines[i - 1];
            this.lines[i - 1] = temp;
        }

        // Modify selection indecies.
        const movedChars = this.lines[selectionEndLine].length + 1;
        this.selectionStart -= movedChars;
        this.selectionEnd -= movedChars;
    }

    // Moves the currently selected lines down by one.
    public moveDown() {
        const selectionStartLine = this.lineByPos(this.selectionStart);
        const selectionEndLine = this.lineByPos(this.selectionEnd);
        if (selectionEndLine >= this.lines.length) return;

        // Swap lines in place till the entire selection is moved.
        for (let i = selectionEndLine; i >= selectionStartLine; i--) {
            const temp = this.lines[i];
            this.lines[i] = this.lines[i + 1];
            this.lines[i + 1] = temp;
        }

        // Modify selection indecies.
        const movedChars = this.lines[selectionStartLine].length + 1;
        this.selectionStart += movedChars;
        this.selectionEnd += movedChars;
    }

    // Return a joined view of the file.
    public getContent(): string {
        return this.lines.join("\n");
    }

    // Return the current selection start index.
    public getSelectionStart(): number {
        return this.selectionStart;
    }

    // Return the current selection end index.
    public getSelectionEnd(): number {
        return this.selectionEnd;
    }
}

export class LineEditor {
    private readonly INDENT = "    ";
    private readonly INDENT_LENGTH = this.INDENT.length;

    private lines: string[];
    private selectionStart: number;
    private selectionEnd: number;
    private selectionStartLine: number;
    private selectionEndLine: number;

    constructor(content: string, selectionStart: number, selectionEnd: number) {
        this.lines = content.split("\n");
        this.selectionStart = selectionStart;
        this.selectionEnd = selectionEnd;
        this.selectionStartLine = this.lineByIndex(selectionStart);
        this.selectionEndLine = this.lineByIndex(selectionEnd);
    }

    // Computes the line number at which the index is found.
    // Line breaks are counted as a single character.
    private lineByIndex(index: number): number {
        let charCount = 0;
        let lineCount = 0;
        for (const line of this.lines) {
            charCount += line.length + 1;
            if (charCount > index) {
                return lineCount;
            }
            lineCount++;
        }
        throw new Error("LineEditor: invalid cursor position");
    }

    private startByLine(index: number): number {
        let charCount = 0;
        for (let i = 0; i < index; i++) {
            charCount += this.lines[i].length;
        }
        return charCount;
    }

    public indent() {
        // Modify line contents.
        for (let i = this.selectionStartLine; i <= this.selectionEndLine; i++) {
            this.lines[i] = this.INDENT + this.lines[i];
        }

        // Modify selection indecies.
        this.selectionStart += this.INDENT_LENGTH;
        const editedLines = this.selectionEndLine - this.selectionStartLine + 1;
        this.selectionEnd += this.INDENT_LENGTH * editedLines;
    }

    public unindent() {
        // Modify line contents (if they start with an indent).
        let editedLines = 0;
        let firstLineEdited = false;
        for (let i = this.selectionStartLine; i <= this.selectionEndLine; i++) {
            if (this.lines[i].startsWith(this.INDENT)) {
                if (i === this.selectionStartLine) firstLineEdited = true;
                this.lines[i] = this.lines[i].substr(this.INDENT_LENGTH);
                editedLines++;
            }
        }

        // Modify selection indecies.
        if (firstLineEdited) {
            this.selectionStart -= this.INDENT_LENGTH;
        }
        this.selectionEnd -= this.INDENT_LENGTH * editedLines;

        // Keep cursor on same line it was before unindent.
        if (this.selectionStartLine !== this.lineByIndex(this.selectionStart)) {
            this.selectionStart = this.startByLine(this.selectionStartLine);
        }
        if (this.selectionEndLine !== this.lineByIndex(this.selectionEnd)) {
            this.selectionEnd = this.startByLine(this.selectionEndLine);
        }
    }

    // Join the lines into a contiguous string.
    public toString(): string {
        return this.lines.join("\n");
    }

    public getSelectionStart(): number {
        return this.selectionStart;
    }

    public getSelectionEnd(): number {
        return this.selectionEnd;
    }
}

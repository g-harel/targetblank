const INDENT = "    ";
const INDENT_LENGTH = INDENT.length;
const COMMENT = "#";
const COMMENT_LENGTH = COMMENT.length;
const COMMENT_REPLACE_PATTERN = /^(\s*)# ?/g;

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

    // Decides whether or not a line contains content.
    private lineIsEmpty(line: number): boolean {
        return this.lines[line].trim() === "";
    }

    // Calculates the number of spaces at the start of the line.
    private leadingSpace(line: number): number {
        const contents = this.lines[line];
        let count = 0;
        while (contents[count] === " ") count++;
        return count;
    }

    // Increase indentation level of all selected lines.
    // Cursor positions are updated.
    public indent() {
        const selectionStartLine = this.lineByPos(this.selectionStart);
        const selectionEndLine = this.lineByPos(this.selectionEnd);

        // Modify line contents.
        let firstLineAdded = 0;
        let totalAdded = 0;
        for (let i = selectionStartLine; i <= selectionEndLine; i++) {
            if (this.lineIsEmpty(i)) continue;
            const add = INDENT_LENGTH - (this.leadingSpace(i) % INDENT_LENGTH);
            this.lines[i] = INDENT.slice(0, add) + this.lines[i];
            if (i === selectionStartLine) firstLineAdded = add;
            totalAdded += add;
        }

        // Modify selection indecies.
        this.selectionStart += firstLineAdded;
        this.selectionEnd += totalAdded;
    }

    // Decrease indentation level of all selected lines.
    // Cursor positions are updated, but will not move across lines.
    public unindent() {
        const selectionStartLine = this.lineByPos(this.selectionStart);
        const selectionEndLine = this.lineByPos(this.selectionEnd);

        // Modify line contents (if they start with an indent).
        let firstLineRemoved = 0;
        let totalRemoved = 0;
        for (let i = selectionStartLine; i <= selectionEndLine; i++) {
            if (this.lineIsEmpty(i)) continue;
            const whitespace = this.leadingSpace(i);
            if (whitespace === 0) continue;
            let remove = whitespace % INDENT_LENGTH;
            if (remove === 0) remove = INDENT_LENGTH;
            this.lines[i] = this.lines[i].substr(remove);
            if (i === selectionStartLine) firstLineRemoved = remove;
            totalRemoved += remove;
        }

        // Modify selection indecies.
        this.selectionStart -= firstLineRemoved;
        this.selectionEnd -= totalRemoved;

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

    public toggleComment() {
        const selectionStartLine = this.lineByPos(this.selectionStart);
        const selectionEndLine = this.lineByPos(this.selectionEnd);

        // Decide whether to comment or un-comment selection. Lines will be
        // commented out if any of the selected lines is not currently commented.
        let comment = false;
        for (let i = selectionStartLine; i <= selectionEndLine; i++) {
            if (this.lineIsEmpty(i)) continue;
            comment = comment || !this.lines[i].trim().startsWith(COMMENT);
        }

        if (comment) {
            const startIndexInLine =
                this.selectionStart - this.posByLine(selectionStartLine);
            const endIndexInLine =
                this.selectionEnd - this.posByLine(selectionEndLine);
            // Find highest possible level that can be commented.
            let level = Infinity;
            for (let i = selectionStartLine; i <= selectionEndLine; i++) {
                if (this.lineIsEmpty(i)) continue;
                level = Math.min(level, this.leadingSpace(i));
            }
            // Only comment at valid indentation levels.
            level = level - (level % INDENT_LENGTH);
            let totalAdded = 0;
            for (let i = selectionStartLine; i <= selectionEndLine; i++) {
                if (this.lineIsEmpty(i)) continue;
                const line = this.lines[i];
                this.lines[i] = `${line.slice(0, level)}${COMMENT} ${line.slice(
                    level,
                )}`;
                totalAdded += COMMENT_LENGTH + 1;
            }
            // Modify selection indecies.
            if (startIndexInLine >= level) {
                this.selectionStart += COMMENT_LENGTH + 1;
            }
            if (endIndexInLine >= level) {
                this.selectionEnd += totalAdded;
            } else {
                this.selectionEnd += totalAdded - COMMENT_LENGTH - 1;
            }
        } else {
            const startIndexInLine =
                this.selectionStart - this.posByLine(selectionStartLine);
            const endIndexInLine =
                this.selectionEnd - this.posByLine(selectionEndLine);
            let firstLineRemoved = 0;
            let firstLineRemovedIndex = 0;
            let lastLineRemoved = 0;
            let lastLineRemovedIndex = 0;
            let totalRemoved = 0;
            for (let i = selectionStartLine; i <= selectionEndLine; i++) {
                if (this.lineIsEmpty(i)) continue;
                if (i === selectionStartLine) {
                    firstLineRemovedIndex = this.leadingSpace(i);
                }
                if (i === selectionEndLine - 1) {
                    lastLineRemovedIndex = this.leadingSpace(i);
                }
                const line = this.lines[i];
                this.lines[i] = line.replace(COMMENT_REPLACE_PATTERN, "$1");
                const removed = line.length - this.lines[i].length;
                totalRemoved += removed;
                if (i === selectionStartLine) firstLineRemoved = removed;
                if (i === selectionEndLine - 1) lastLineRemoved = removed;
            }
            // Modify selection indecies.
            if (startIndexInLine >= firstLineRemovedIndex) {
                this.selectionStart -= firstLineRemoved;
            }
            if (endIndexInLine >= lastLineRemovedIndex) {
                this.selectionEnd -= totalRemoved;
            } else {
                this.selectionEnd -= totalRemoved - lastLineRemoved;
            }
        }
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

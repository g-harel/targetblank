import {FileEditor} from "./file";

describe("FileEditor", () => {
    it("should not modify contents on creation", () => {
        const file = "\nabc\n  abc\n  1 2 3";
        const fileEditor = new FileEditor(file, 0, 0);
        expect(fileEditor.getContent()).toBe(file);
    });

    it("should reject negative selection start index", () => {
        const file = "";
        expect(() => new FileEditor(file, -1, 0)).toThrow();
    });

    it("should reject negative selection end index", () => {
        const file = "";
        expect(() => new FileEditor(file, 0, -1)).toThrow();
    });

    it("should reject out of range selection start index", () => {
        const file = "";
        expect(() => new FileEditor(file, 10, 0)).toThrow();
    });

    it("should reject out of range selection end index", () => {
        const file = "";
        expect(() => new FileEditor(file, 0, 10)).toThrow();
    });

    describe("indent", () => {
        it("should work correctly on a single line", () => {
            const input = "abc";
            const cursor = 0;
            const expected = "    abc";

            const fileEditor = new FileEditor(input, cursor, cursor);
            fileEditor.indent();

            expect(fileEditor.getContent()).toBe(expected);
        });

        it("should work correctly on multiple lines", () => {
            const input = "abc\n    123\n456\n    xyz";
            const cursorStart = 6;
            const cursorEnd = 14;
            const expected = "abc\n        123\n    456\n    xyz";

            const fileEditor = new FileEditor(input, cursorStart, cursorEnd);
            fileEditor.indent();

            expect(fileEditor.getContent()).toBe(expected);
        });

        it("should correctly adjust cursor position", () => {
            const input = "  123";
            const cursor = 2;
            const expectedCursor = 4;

            const fileEditor = new FileEditor(input, cursor, cursor);
            fileEditor.indent();

            expect(fileEditor.getSelectionStart()).toBe(expectedCursor);
            expect(fileEditor.getSelectionEnd()).toBe(expectedCursor);
        });

        it("should correctly adjust cursor position on multiple lines", () => {
            const input = "  abc\n  123\n 456\nxyz";
            const cursorStart = 10;
            const cursorEnd = 18;
            const expectedStart = 12;
            const expectedEnd = 27;

            const fileEditor = new FileEditor(input, cursorStart, cursorEnd);
            fileEditor.indent();

            expect(fileEditor.getSelectionStart()).toBe(expectedStart);
            expect(fileEditor.getSelectionEnd()).toBe(expectedEnd);
        });
    });

    describe("unindent", () => {
        it("should work correctly on a single line", () => {
            const input = "    abc";
            const cursor = 0;
            const expected = "abc";

            const fileEditor = new FileEditor(input, cursor, cursor);
            fileEditor.unindent();

            expect(fileEditor.getContent()).toBe(expected);
        });

        it("should be a noop if the line is not indented", () => {
            const input = "abc";
            const cursor = 0;
            const expected = "abc";

            const fileEditor = new FileEditor(input, cursor, cursor);
            fileEditor.unindent();

            expect(fileEditor.getContent()).toBe(expected);
        });

        it("should work correctly on partially indented lines", () => {
            const input = "  abc";
            const cursor = 0;
            const expected = "abc";

            const fileEditor = new FileEditor(input, cursor, cursor);
            fileEditor.unindent();

            expect(fileEditor.getContent()).toBe(expected);
        });

        it("should work correctly on multiple lines", () => {
            const input = "abc\n    123\n456\n    xyz";
            const cursorStart = 6;
            const cursorEnd = 14;
            const expected = "abc\n123\n456\n    xyz";

            const fileEditor = new FileEditor(input, cursorStart, cursorEnd);
            fileEditor.unindent();

            expect(fileEditor.getContent()).toBe(expected);
        });

        it("should correctly adjust cursor position", () => {
            const input = "        123";
            const cursor = 7;
            const expectedCursor = 3;

            const fileEditor = new FileEditor(input, cursor, cursor);
            fileEditor.unindent();

            expect(fileEditor.getSelectionStart()).toBe(expectedCursor);
            expect(fileEditor.getSelectionEnd()).toBe(expectedCursor);
        });

        it("should correctly adjust cursor position on multiple lines", () => {
            const input = "  abc\n    123\n  456\n      xyz";
            const cursorStart = 10;
            const cursorEnd = 25;
            const expectedStart = 6;
            const expectedEnd = 17;

            const fileEditor = new FileEditor(input, cursorStart, cursorEnd);
            fileEditor.unindent();

            expect(fileEditor.getSelectionStart()).toBe(expectedStart);
            expect(fileEditor.getSelectionEnd()).toBe(expectedEnd);
        });

        it("should not change the cursor's line when it close to the start", () => {
            const input = "  abc\n    123\n  456\n      xyz";
            const cursorStart = 6;
            const cursorEnd = 21;
            const expectedStart = 6;
            const expectedEnd = 14;

            const fileEditor = new FileEditor(input, cursorStart, cursorEnd);
            fileEditor.unindent();

            expect(fileEditor.getSelectionStart()).toBe(expectedStart);
            expect(fileEditor.getSelectionEnd()).toBe(expectedEnd);
        });
    });
});
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
            const expectedCursor = 6;

            const fileEditor = new FileEditor(input, cursor, cursor);
            fileEditor.indent();

            expect(fileEditor.getSelectionStart()).toBe(expectedCursor);
            expect(fileEditor.getSelectionEnd()).toBe(expectedCursor);
        });

        it("should correctly adjust cursor position", () => {
            const input = "  abc\n  123\n 456\nxyz";
            const cursorStart = 10;
            const cursorEnd = 13;
            const expectedStart = 14;
            const expectedEnd = 21;

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

        it("should be a noop if the line is partially indented", () => {
            const input = "  abc";
            const cursor = 0;
            const expected = "  abc";

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
    });
});

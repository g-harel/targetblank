import {moveUp, moveDown} from "./movement";

describe("website/internal/editor/movement", () => {
    describe("moveUp", () => {
        it("should work correctly on a single line", () => {
            const value = "a\nbc";
            const selectionStart = 2;
            const selectionEnd = selectionStart;
            const expected = "bc\na";

            const temp = moveUp({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should work correctly on multiple lines", () => {
            const value = "a\n\nabc";
            const selectionStart = 2;
            const selectionEnd = 4;
            const expected = "\nabc\na";

            const temp = moveUp({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should correctly adjust cursor position", () => {
            const value = "\n123\n\n456";
            const selectionStart = 8;
            const selectionEnd = selectionStart;
            const expectedCursor = 7;

            const temp = moveUp({value, selectionStart, selectionEnd});

            expect(temp.selectionStart).toBe(expectedCursor);
            expect(temp.selectionEnd).toBe(expectedCursor);
        });

        it("should correctly adjust cursor position on multiple lines", () => {
            const value = "123\n    456\n\n789";
            const selectionStart = 6;
            const selectionEnd = 16;
            const expectedCursorStart = 2;
            const expectedCursorEnd = 12;

            const temp = moveUp({value, selectionStart, selectionEnd});

            expect(temp.selectionStart).toBe(expectedCursorStart);
            expect(temp.selectionEnd).toBe(expectedCursorEnd);
        });

        it("should be a noop on the first line", () => {
            const value = "xyz\n123";
            const selectionStart = 1;
            const selectionEnd = 2;
            const expected = value;
            const expectedStart = selectionStart;
            const expectedEnd = selectionEnd;

            const temp = moveUp({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
            expect(temp.selectionStart).toBe(expectedStart);
            expect(temp.selectionEnd).toBe(expectedEnd);
        });
    });

    describe("moveDown", () => {
        it("should work correctly on a single line", () => {
            const value = "a\nbc";
            const selectionStart = 1;
            const selectionEnd = selectionStart;
            const expected = "bc\na";

            const temp = moveDown({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should work correctly on multiple lines", () => {
            const value = "a\n\nabc";
            const selectionStart = 0;
            const selectionEnd = 2;
            const expected = "abc\na\n";

            const temp = moveDown({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should correctly adjust cursor position", () => {
            const value = "\n123\n\n456";
            const selectionStart = 5;
            const selectionEnd = selectionStart;
            const expectedCursor = 9;

            const temp = moveDown({value, selectionStart, selectionEnd});

            expect(temp.selectionStart).toBe(expectedCursor);
            expect(temp.selectionEnd).toBe(expectedCursor);
        });

        it("should correctly adjust cursor position on multiple lines", () => {
            const value = "123\n    456\n\n789";
            const selectionStart = 6;
            const selectionEnd = 12;
            const expectedCursorStart = 10;
            const expectedCursorEnd = 16;

            const temp = moveDown({value, selectionStart, selectionEnd});

            expect(temp.selectionStart).toBe(expectedCursorStart);
            expect(temp.selectionEnd).toBe(expectedCursorEnd);
        });

        it("should be a noop on the last line", () => {
            const value = "xyz\n123";
            const selectionStart = 1;
            const selectionEnd = 5;
            const expected = value;
            const expectedStart = selectionStart;
            const expectedEnd = selectionEnd;

            const temp = moveDown({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
            expect(temp.selectionStart).toBe(expectedStart);
            expect(temp.selectionEnd).toBe(expectedEnd);
        });
    });
});

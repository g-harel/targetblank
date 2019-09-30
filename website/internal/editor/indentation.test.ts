import {indent, unindent} from "./indentation";

describe("website/internal/editor/indentation", () => {
    describe("indent", () => {
        it("should work correctly on a single line", () => {
            const value = "abc";
            const selectionStart = 0;
            const selectionEnd = selectionStart;
            const expected = "    abc";

            const temp = indent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should work correctly on multiple lines", () => {
            const value = "abc\n    123\n456\n    xyz";
            const selectionStart = 6;
            const selectionEnd = 14;
            const expected = "abc\n        123\n    456\n    xyz";

            const temp = indent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should correctly adjust cursor position", () => {
            const value = "  123";
            const selectionStart = 2;
            const selectionEnd = selectionStart;
            const expectedCursor = 4;

            const temp = indent({value, selectionStart, selectionEnd});

            expect(temp.selectionStart).toBe(expectedCursor);
            expect(temp.selectionEnd).toBe(expectedCursor);
        });

        it("should correctly adjust cursor position on multiple lines", () => {
            const value = "  abc\n  123\n 456\nxyz";
            const selectionStart = 10;
            const selectionEnd = 18;
            const expectedStart = 12;
            const expectedEnd = 27;

            const temp = indent({value, selectionStart, selectionEnd});

            expect(temp.selectionStart).toBe(expectedStart);
            expect(temp.selectionEnd).toBe(expectedEnd);
        });
    });

    describe("unindent", () => {
        it("should work correctly on a single line", () => {
            const value = "    abc";
            const selectionStart = 0;
            const selectionEnd = selectionStart;
            const expected = "abc";

            const temp = unindent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should be a noop if the line is not indented", () => {
            const value = "abc";
            const selectionStart = 0;
            const selectionEnd = selectionStart;
            const expected = "abc";

            const temp = unindent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should work correctly on partially indented lines", () => {
            const value = "  abc";
            const selectionStart = 0;
            const selectionEnd = selectionStart;
            const expected = "abc";

            const temp = unindent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should work correctly on multiple lines", () => {
            const value = "abc\n    123\n456\n    xyz";
            const selectionStart = 6;
            const selectionEnd = 14;
            const expected = "abc\n123\n456\n    xyz";

            const temp = unindent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should correctly adjust cursor position", () => {
            const value = "        123";
            const selectionStart = 7;
            const selectionEnd = selectionStart;
            const expectedCursor = 3;

            const temp = unindent({value, selectionStart, selectionEnd});

            expect(temp.selectionStart).toBe(expectedCursor);
            expect(temp.selectionEnd).toBe(expectedCursor);
        });

        it("should correctly adjust cursor position on multiple lines", () => {
            const value = "  abc\n    123\n  456\n      xyz";
            const selectionStart = 10;
            const selectionEnd = 25;
            const expectedStart = 6;
            const expectedEnd = 17;

            const temp = unindent({value, selectionStart, selectionEnd});

            expect(temp.selectionStart).toBe(expectedStart);
            expect(temp.selectionEnd).toBe(expectedEnd);
        });

        it("should not change the cursor's line when it close to the start", () => {
            const value = "  abc\n    123\n  456\n      xyz";
            const selectionStart = 6;
            const selectionEnd = 21;
            const expectedStart = 6;
            const expectedEnd = 14;

            const temp = unindent({value, selectionStart, selectionEnd});

            expect(temp.selectionStart).toBe(expectedStart);
            expect(temp.selectionEnd).toBe(expectedEnd);
        });
    });
});
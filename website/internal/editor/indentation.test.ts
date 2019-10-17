import {indent, unindent, newline} from "./indentation";
import {genRandomState, expectEqual} from "./testing";

describe("website/internal/editor/indentation", () => {
    describe("indent", () => {
        it("should correctly add indentation on a single line", () => {
            const value = "abc";
            const selectionStart = 0;
            const selectionEnd = selectionStart;
            const expected = "    abc";

            const temp = indent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should add indentation on a single empty line", () => {
            const value = "";
            const selectionStart = 0;
            const selectionEnd = selectionStart;
            const expected = "    ";

            const temp = indent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should snap to the nearest higher indentation level", () => {
            const value = "  abc";
            const selectionStart = 0;
            const selectionEnd = selectionStart;
            const expected = "    abc";

            const temp = indent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should indent all lines when multiple selected", () => {
            const value = "abc\n    123\n456\n    xyz";
            const selectionStart = 6;
            const selectionEnd = 14;
            const expected = "abc\n        123\n    456\n    xyz";

            const temp = indent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should not indent empty lines when multiple selected", () => {
            const value = "abc\n\n123";
            const selectionStart = 2;
            const selectionEnd = 6;
            const expected = "    abc\n\n    123";

            const temp = indent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should adjust cursor position", () => {
            const value = "  123";
            const selectionStart = 2;
            const selectionEnd = selectionStart;
            const expectedCursor = 4;

            const temp = indent({value, selectionStart, selectionEnd});

            expect(temp.selectionStart).toBe(expectedCursor);
            expect(temp.selectionEnd).toBe(expectedCursor);
        });

        it("should adjust cursor position when multiple selected", () => {
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
        it("should correctly remove indentation from single line", () => {
            const value = "    abc";
            const selectionStart = 0;
            const selectionEnd = selectionStart;
            const expected = "abc";

            const temp = unindent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should make not changes if line is not indented", () => {
            const value = "abc";
            const selectionStart = 0;
            const selectionEnd = selectionStart;
            const expected = value;

            const temp = unindent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should snap to the nearest lower indentation level", () => {
            const value = "      abc";
            const selectionStart = 0;
            const selectionEnd = selectionStart;
            const expected = "    abc";

            const temp = unindent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should un-indent all selected lines", () => {
            const value = "abc\n    123\n456\n    xyz";
            const selectionStart = 6;
            const selectionEnd = 14;
            const expected = "abc\n123\n456\n    xyz";

            const temp = unindent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should ignore empty lines when un-indenting", () => {
            const value = "    abc\n\n        123";
            const selectionStart = 3;
            const selectionEnd = 19;
            const expected = "abc\n\n    123";

            const temp = unindent({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should adjust cursor position", () => {
            const value = "        123";
            const selectionStart = 7;
            const selectionEnd = selectionStart;
            const expectedCursor = 3;

            const temp = unindent({value, selectionStart, selectionEnd});

            expect(temp.selectionStart).toBe(expectedCursor);
            expect(temp.selectionEnd).toBe(expectedCursor);
        });

        it("should adjust cursor position when multiple selected", () => {
            const value = "  abc\n    123\n  456\n      xyz";
            const selectionStart = 10;
            const selectionEnd = 25;
            const expectedStart = 6;
            const expectedEnd = 17;

            const temp = unindent({value, selectionStart, selectionEnd});

            expect(temp.selectionStart).toBe(expectedStart);
            expect(temp.selectionEnd).toBe(expectedEnd);
        });

        it("should not change the cursor's line when it bottoms out", () => {
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

    // Checks that randomly generated states return to the initial state
    // after being cycled (unindent + indent).
    it("should be stable", () => {
        for (let i = 0; i < 32; i++) {
            // Initial state is pre-indented to normalize indentation levels
            // and avoid bottoming out the line with unindent.
            const initialEditorState = indent(genRandomState(16));
            const indentedEditorState = indent(initialEditorState);
            expectEqual(unindent(indentedEditorState), initialEditorState);
        }
    });

    describe("newline", () => {
        it("should add a new line", () => {
            const value = "abc";
            const selectionStart = 3;
            const selectionEnd = selectionStart;
            const expected = "abc\n";

            const temp = newline({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should keep the same indentation", () => {
            const value = "   abc";
            const selectionStart = 6;
            const selectionEnd = selectionStart;
            const expected = "   abc\n   ";

            const temp = newline({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should keep the same indentation when cursor in whitespace", () => {
            const value = "     123";
            const selectionStart = 3;
            const selectionEnd = selectionStart;
            const expected = "   \n     123";

            const temp = newline({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should split up lines when cursor in content", () => {
            const value = "  123";
            const selectionStart = 4;
            const selectionEnd = selectionStart;
            const expected = "  12\n  3";

            const temp = newline({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should add new line after the end of the selection", () => {
            const value = "abc\n    123\n 456\nxyz";
            const selectionStart = 10;
            const selectionEnd = 14;
            const expected = "abc\n    123\n 4\n 56\nxyz";

            const temp = newline({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should adjust cursor position", () => {
            const value = "    123";
            const selectionStart = 4;
            const selectionEnd = selectionStart;
            const expectedCursor = 9;

            const temp = newline({value, selectionStart, selectionEnd});

            expect(temp.selectionStart).toBe(expectedCursor);
            expect(temp.selectionEnd).toBe(expectedCursor);
        });

        it("should adjust cursor position when multiple selected", () => {
            const value = "  abc\n    123\n  456\n      xyz";
            const selectionStart = 10;
            const selectionEnd = 25;
            const expectedStart = 31;
            const expectedEnd = expectedStart;

            const temp = newline({value, selectionStart, selectionEnd});

            expect(temp.selectionStart).toBe(expectedStart);
            expect(temp.selectionEnd).toBe(expectedEnd);
        });
    });
});

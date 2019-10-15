import {toggleComment} from "./comment";
import {EditorState} from ".";

describe("website/internal/editor/comment", () => {
    describe("toggleComment", () => {
        it("should comment non-commented lines", () => {
            const value = "abc";
            const selectionStart = 0;
            const selectionEnd = selectionStart;
            const expected = "# abc";

            const temp = toggleComment({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should un-comment commented lines", () => {
            const value = "# abc";
            const selectionStart = 0;
            const selectionEnd = selectionStart;
            const expected = "abc";

            const temp = toggleComment({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should comment multiple non-commented lines", () => {
            const value = "123\n456\n789";
            const selectionStart = 6;
            const selectionEnd = 10;
            const expected = "123\n# 456\n# 789";

            const temp = toggleComment({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should un-comment multiple commented lines", () => {
            const value = "# 123\n# 456\n# 789";
            const selectionStart = 9;
            const selectionEnd = 15;
            const expected = "# 123\n456\n789";

            const temp = toggleComment({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should comment multiple mixed lines", () => {
            const value = "# 123\n456\n# 789";
            const selectionStart = 3;
            const selectionEnd = 13;
            const expected = "# # 123\n# 456\n# # 789";

            const temp = toggleComment({value, selectionStart, selectionEnd});

            expect(temp.value).toBe(expected);
        });

        it("should ignore empty lines when toggling multiple", () => {
            const value = "# 123\n\n# 456";
            const selectionStart = 3;
            const selectionEnd = 10;
            const expected = "123\n\n456";

            const temp = toggleComment({
                value,
                selectionStart,
                selectionEnd,
            });

            expect(temp.value).toBe(expected);
        });

        describe("commentOn", () => {
            it("should not comment empty lines when multiple selected", () => {
                const value = "123\n\n456";
                const selectionStart = 2;
                const selectionEnd = 6;
                const expected = "# 123\n\n# 456";

                const temp = toggleComment({
                    value,
                    selectionStart,
                    selectionEnd,
                });

                expect(temp.value).toBe(expected);
            });

            it("should add comment at highest possible indentation level", () => {
                const value = "        123\n\n    456";
                const selectionStart = 9;
                const selectionEnd = 16;
                const expected = "    #     123\n\n    # 456";

                const temp = toggleComment({
                    value,
                    selectionStart,
                    selectionEnd,
                });

                expect(temp.value).toBe(expected);
            });

            it("should only comment at valid indentation levels", () => {
                const value = "   123\n\n    456";
                const selectionStart = 1;
                const selectionEnd = 12;
                const expected = "#    123\n\n#     456";

                const temp = toggleComment({
                    value,
                    selectionStart,
                    selectionEnd,
                });

                expect(temp.value).toBe(expected);
            });

            it("should adjust cursor position when after comment", () => {
                const value = "    123";
                const selectionStart = 5;
                const selectionEnd = selectionStart;
                const expectedCursor = 7;

                const temp = toggleComment({
                    value,
                    selectionStart,
                    selectionEnd,
                });

                expect(temp.selectionStart).toBe(expectedCursor);
                expect(temp.selectionEnd).toBe(expectedCursor);
            });

            it("should not adjust cursor position when before comment", () => {
                const value = "    123";
                const selectionStart = 2;
                const selectionEnd = selectionStart;
                const expectedCursor = 2;

                const temp = toggleComment({
                    value,
                    selectionStart,
                    selectionEnd,
                });

                expect(temp.selectionStart).toBe(expectedCursor);
                expect(temp.selectionEnd).toBe(expectedCursor);
            });

            it("should correctly adjust cursor position when multiple selected", () => {
                const value = "    abc\n    123\n    456\n      xyz";
                const selectionStart = 5;
                const selectionEnd = 27;
                const expectedStart = 7;
                const expectedEnd = 33;

                const temp = toggleComment({
                    value,
                    selectionStart,
                    selectionEnd,
                });

                expect(temp.selectionStart).toBe(expectedStart);
                expect(temp.selectionEnd).toBe(expectedEnd);
            });
        });

        describe("commentOff", () => {
            it("should remove comments at any indentation level", () => {
                const value = "       # 123\n\n  # 456";
                const selectionStart = 9;
                const selectionEnd = 16;
                const expected = "       123\n\n  456";

                const temp = toggleComment({
                    value,
                    selectionStart,
                    selectionEnd,
                });

                expect(temp.value).toBe(expected);
            });

            it("should adjust cursor position when cursor after comment", () => {
                const value = "   # 123";
                const selectionStart = 5;
                const selectionEnd = selectionStart;
                const expectedCursor = 3;

                const temp = toggleComment({
                    value,
                    selectionStart,
                    selectionEnd,
                });

                expect(temp.selectionStart).toBe(expectedCursor);
                expect(temp.selectionEnd).toBe(expectedCursor);
            });

            it("should not adjust cursor position when before comment", () => {
                const value = "    # 123";
                const selectionStart = 1;
                const selectionEnd = selectionStart;
                const expectedCursor = 1;

                const temp = toggleComment({
                    value,
                    selectionStart,
                    selectionEnd,
                });

                expect(temp.selectionStart).toBe(expectedCursor);
                expect(temp.selectionEnd).toBe(expectedCursor);
            });

            it("should correctly adjust cursor position when multiple selected", () => {
                const value = "    # abc\n    # 123\n    # 456\n      # xyz";
                const selectionStart = 7;
                const selectionEnd = 32;
                const expectedStart = 5;
                const expectedEnd = 26;

                const temp = toggleComment({
                    value,
                    selectionStart,
                    selectionEnd,
                });

                expect(temp.selectionStart).toBe(expectedStart);
                expect(temp.selectionEnd).toBe(expectedEnd);
            });
        });

        // Checks that randomly generated states return to the initial state
        // after being toggled twice (commented + un-commented).
        for (let i = 0; i < 32; i++) {
            it(`it should be stable #${i}`, () => {
                const lineCount = Math.ceil(16 * Math.random());
                const lines = [];
                for (let j = 0; j < lineCount; j++) {
                    const whitespaceCount = Math.floor(16 * Math.random());
                    const contentCount = Math.ceil(16 * Math.random());
                    lines.push(
                        " ".repeat(whitespaceCount) + "a".repeat(contentCount),
                    );
                }
                const value = lines.join("\n");
                const a = Math.floor(value.length * Math.random());
                const b = Math.floor(value.length * Math.random());

                const initialEditorState: EditorState = {
                    value,
                    selectionStart: Math.min(a, b),
                    selectionEnd: Math.max(a, b),
                };
                const commentedEditorState = toggleComment(initialEditorState);
                expect(toggleComment(commentedEditorState)).toEqual(initialEditorState);
            });
        }
    });
});

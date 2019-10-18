import {indent, unindent, newline} from "./indentation";
import {genRandomState, expectCommand, expectEqual} from "./testing";

describe("website/internal/editor/indentation", () => {
    describe("indent", () => {
        it("should correctly add indentation on a single line", () => {
            expectCommand(indent)
                .withValue("abc")
                .withCursor(0)
                .toProduceValue("    abc");
        });

        it("should add indentation on a single empty line", () => {
            expectCommand(indent)
                .withValue("")
                .withCursor(0)
                .toProduceValue("    ");
        });

        it("should snap to the nearest higher indentation level", () => {
            expectCommand(indent)
                .withValue("  abc")
                .withCursor(0)
                .toProduceValue("    abc");
        });

        it("should indent all lines when multiple selected", () => {
            expectCommand(indent)
                .withValue("abc\n    123\n456\n    xyz")
                .withSelection(6, 14)
                .toProduceValue("abc\n        123\n    456\n    xyz");
        });

        it("should not indent empty lines when multiple selected", () => {
            expectCommand(indent)
                .withValue("abc\n\n123")
                .withSelection(2, 6)
                .toProduceValue("    abc\n\n    123");
        });

        it("should adjust cursor position", () => {
            expectCommand(indent)
                .withValue("  123")
                .withCursor(2)
                .toProduceCursor(4);
        });

        it("should adjust cursor position when multiple selected", () => {
            expectCommand(indent)
                .withValue("  abc\n  123\n 456\nxyz")
                .withSelection(10, 18)
                .toProduceSelection(12, 27);
        });
    });

    describe("unindent", () => {
        it("should correctly remove indentation from single line", () => {
            expectCommand(unindent)
                .withValue("    abc")
                .withCursor(0)
                .toProduceValue("abc");
        });

        it("should make not changes if line is not indented", () => {
            expectCommand(unindent)
                .withValue("abc")
                .withCursor(0)
                .toProduceUnchanged();
        });

        it("should snap to the nearest lower indentation level", () => {
            expectCommand(unindent)
                .withValue("       abc")
                .withCursor(0)
                .toProduceValue("    abc");
        });

        it("should un-indent all selected lines", () => {
            expectCommand(unindent)
                .withValue("abc\n    123\n456\n    xyz")
                .withSelection(6, 14)
                .toProduceValue("abc\n123\n456\n    xyz");
        });

        it("should ignore empty lines when un-indenting", () => {
            expectCommand(unindent)
                .withValue("    abc\n\n        123")
                .withSelection(3, 19)
                .toProduceValue("abc\n\n    123");
        });

        it("should adjust cursor position", () => {
            expectCommand(unindent)
                .withValue("        123")
                .withCursor(7)
                .toProduceCursor(3);
        });

        it("should adjust cursor position when multiple selected", () => {
            expectCommand(unindent)
                .withValue("  abc\n    123\n  456\n      xyz")
                .withSelection(10, 25)
                .toProduceSelection(6, 17);
        });

        it("should not change the cursor's line when it bottoms out", () => {
            expectCommand(unindent)
                .withValue("  abc\n    123\n  456\n      xyz")
                .withSelection(6, 21)
                .toProduceSelection(6, 14);
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
            expectCommand(newline)
                .withValue("abc")
                .withCursor(3)
                .toProduceValue("abc\n");
        });

        it("should keep the same indentation", () => {
            expectCommand(newline)
                .withValue("   abc")
                .withCursor(6)
                .toProduceValue("   abc\n   ");
        });

        it("should keep the same indentation when cursor in whitespace", () => {
            expectCommand(newline)
                .withValue("     123")
                .withCursor(3)
                .toProduceValue("   \n     123");
        });

        it("should split up lines when cursor in content", () => {
            expectCommand(newline)
                .withValue("  123")
                .withCursor(4)
                .toProduceValue("  12\n  3");
        });

        it("should add new line after the end of the selection", () => {
            expectCommand(newline)
                .withValue("abc\n    123\n 456\nxyz")
                .withSelection(10, 14)
                .toProduceValue("abc\n    123\n 4\n 56\nxyz");
        });

        it("should adjust cursor position", () => {
            expectCommand(newline)
                .withValue("    123")
                .withCursor(4)
                .toProduceCursor(9);
        });

        it("should adjust cursor position when multiple selected", () => {
            expectCommand(newline)
                .withValue("  abc\n    123\n  456\n      xyz")
                .withSelection(10, 25)
                .toProduceSelection(31, 31);
        });
    });
});

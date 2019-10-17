import {toggleComment} from "./comment";
import {genRandomState, expectCommand, expectEqual} from "./testing";

describe("website/internal/editor/comment", () => {
    describe("toggleComment", () => {
        it("should comment non-commented lines", () => {
            expectCommand(toggleComment)
                .withValue("abc")
                .withCursor(0)
                .toProduceValue("# abc");
        });

        it("should un-comment commented lines", () => {
            expectCommand(toggleComment)
                .withValue("# abc")
                .withCursor(2)
                .toProduceValue("abc");
        });

        it("should comment multiple non-commented lines", () => {
            expectCommand(toggleComment)
                .withValue("123\n456\n789")
                .withSelection(6, 10)
                .toProduceValue("123\n# 456\n# 789");
        });

        it("should un-comment multiple commented lines", () => {
            expectCommand(toggleComment)
                .withValue("# 123\n# 456\n# 789")
                .withSelection(9, 15)
                .toProduceValue("# 123\n456\n789");
        });

        it("should comment multiple mixed lines", () => {
            expectCommand(toggleComment)
                .withValue("# 123\n456\n# 789")
                .withSelection(3, 13)
                .toProduceValue("# # 123\n# 456\n# # 789");
        });

        it("should ignore empty lines when toggling multiple", () => {
            expectCommand(toggleComment)
                .withValue("# 123\n\n# 456")
                .withSelection(3, 10)
                .toProduceValue("123\n\n456");
        });

        describe("commentOn", () => {
            it("should not comment empty lines when multiple selected", () => {
                expectCommand(toggleComment)
                    .withValue("123\n\n456")
                    .withSelection(2, 6)
                    .toProduceValue("# 123\n\n# 456");
            });

            it("should add comment at highest possible indentation level", () => {
                expectCommand(toggleComment)
                    .withValue("        123\n\n    456")
                    .withSelection(9, 16)
                    .toProduceValue("    #     123\n\n    # 456");
            });

            it("should only comment at valid indentation levels", () => {
                expectCommand(toggleComment)
                    .withValue("   123\n\n    456")
                    .withSelection(1, 12)
                    .toProduceValue("#    123\n\n#     456");
            });

            it("should adjust cursor position when after comment", () => {
                expectCommand(toggleComment)
                    .withValue("    123")
                    .withCursor(5)
                    .toProduceCursor(7);
            });

            it("should not adjust cursor position when before comment", () => {
                expectCommand(toggleComment)
                    .withValue("    123")
                    .withCursor(2)
                    .toProduceCursor(2);
            });

            it("should correctly adjust cursor position when multiple selected", () => {
                expectCommand(toggleComment)
                    .withValue("    abc\n    123\n    456\n      xyz")
                    .withSelection(5, 27)
                    .toProduceSelection(7, 33);
            });
        });

        describe("commentOff", () => {
            it("should remove comments at any indentation level", () => {
                expectCommand(toggleComment)
                    .withValue("       # 123\n\n  # 456")
                    .withSelection(9, 16)
                    .toProduceValue("       123\n\n  456");
            });

            it("should adjust cursor position when cursor after comment", () => {
                expectCommand(toggleComment)
                    .withValue("   # 123")
                    .withCursor(5)
                    .toProduceCursor(3);
            });

            it("should not adjust cursor position when before comment", () => {
                expectCommand(toggleComment)
                    .withValue("   # 123")
                    .withCursor(1)
                    .toProduceCursor(1);
            });

            it("should correctly adjust cursor position when multiple selected", () => {
                expectCommand(toggleComment)
                    .withValue("    # abc\n    # 123\n    # 456\n      # xyz")
                    .withSelection(7, 32)
                    .toProduceSelection(5, 26);
            });
        });

        // Checks that randomly generated states return to the initial state
        // after being toggled twice (commented + un-commented).
        it("should be stable", () => {
            for (let i = 0; i < 32; i++) {
                const initialEditorState = genRandomState(16);
                const commentedEditorState = toggleComment(initialEditorState);
                expectEqual(
                    toggleComment(commentedEditorState),
                    initialEditorState,
                );
            }
        });
    });
});

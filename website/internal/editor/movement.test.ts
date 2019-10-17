import {moveUp, moveDown} from "./movement";
import {genRandomState, expectCommand, expectEqual} from "./testing";

describe("website/internal/editor/movement", () => {
    describe("moveUp", () => {
        it("should move line up", () => {
            expectCommand(moveUp)
                .withValue("a\nbc")
                .withCursor(2)
                .toProduceValue("bc\na");
        });

        it("should move all selected lines", () => {
            expectCommand(moveUp)
                .withValue("a\n\nabc")
                .withSelection(2, 4)
                .toProduceValue("\nabc\na");
        });

        it("should adjust cursor position", () => {
            expectCommand(moveUp)
                .withValue("\n123\n\n456")
                .withCursor(8)
                .toProduceCursor(7);
        });

        it("should adjust cursor position when multiple selected", () => {
            expectCommand(moveUp)
                .withValue("123\n    456\n\n789")
                .withSelection(6, 16)
                .toProduceSelection(2, 12);
        });

        it("should not make any changes when used on the first line", () => {
            expectCommand(moveUp)
                .withValue("xyz\n123")
                .withSelection(1, 2)
                .toProduceUnchanged();
        });
    });

    describe("moveDown", () => {
        it("should move line down", () => {
            expectCommand(moveDown)
                .withValue("a\nbc")
                .withCursor(1)
                .toProduceValue("bc\na");
        });

        it("should move all selected lines", () => {
            expectCommand(moveDown)
                .withValue("a\n\nabc")
                .withSelection(0, 2)
                .toProduceValue("abc\na\n");
        });

        it("should adjust cursor position", () => {
            expectCommand(moveDown)
                .withValue("\n123\n\n456")
                .withCursor(5)
                .toProduceCursor(9);
        });

        it("should adjust cursor position when multiple selected", () => {
            expectCommand(moveDown)
                .withValue("123\n    456\n\n789")
                .withSelection(6, 12)
                .toProduceSelection(10, 16);
        });

        it("should not make any changes when used on the last line", () => {
            expectCommand(moveDown)
                .withValue("xyz\n123")
                .withSelection(1, 5)
                .toProduceUnchanged();
        });
    });

    describe("moveUp/moveDown", () => {
        // Checks that randomly generated states return to the initial state
        // after being cycled (up + down).
        xit("should be stable", () => {
            for (let i = 0; i < 32; i++) {
                // Initial state value is wrapped in newlines so it is always
                // possible to move the selection.
                const initialEditorState = genRandomState(16);
                initialEditorState.value = `\n${initialEditorState.value}\n`;
                const movedEditorState = moveUp(initialEditorState);
                expectEqual(moveDown(movedEditorState), initialEditorState);
            }
        });
    });
});

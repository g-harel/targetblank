import {init, render} from "./util";

describe("website/internal/editor/util", () => {
    describe("init", () => {
        it("should not modify contents on creation", () => {
            const value = "\nabc\n  abc\n  1 2 3";
            const temp = init({value, selectionStart: 0, selectionEnd: 9});
            expect(render(temp).value).toBe(value);
        });

        it("should reject negative selection start index", () => {
            const value = "";
            expect(() =>
                init({value, selectionStart: -1, selectionEnd: 0}),
            ).toThrow();
        });

        it("should reject negative selection end index", () => {
            const value = "";
            expect(() =>
                init({value, selectionStart: 0, selectionEnd: -1}),
            ).toThrow();
        });

        it("should reject out of range selection start index", () => {
            const value = "";
            expect(() =>
                init({value, selectionStart: 10, selectionEnd: 0}),
            ).toThrow();
        });

        it("should reject out of range selection end index", () => {
            const value = "";
            expect(() =>
                init({value, selectionStart: 0, selectionEnd: 10}),
            ).toThrow();
        });

        it("should reject invalid start and end ranges", () => {
            const value = "aaaaaaaaaaaaaaaa";
            expect(() =>
                init({value, selectionStart: 10, selectionEnd: 0}),
            ).toThrow();
        });
    });
});

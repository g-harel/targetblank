import {LineEditor} from "./lines";

describe("LineEditor", () => {
    it("should not modify contents on creation", () => {
        const src = "" + "abc" + "  abc" + "  1 2 3";
        expect(new LineEditor(src, 0, 0).toString()).toBe(src);
    });
});

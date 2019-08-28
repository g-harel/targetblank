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
});

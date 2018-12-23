import okwolo from "okwolo/standard";
import h from "okwolo/src/h";

// app instance singleton
export const app = okwolo();

// okwolo's h function is attached to the global object
(window as any).h = h;

import app from "./app";

app.use("target", document.body);

(window as any).app = app;

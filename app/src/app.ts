import okwolo from "okwolo/lite";

const app = okwolo();

app.setState({});

app("/", () => () => (
    "home"
));

app(/^\/(\w{6})$/g, params => () => (
    params[0]
));

app("**", () => () => (
    "404"
));

export default app;

import okwolo from "okwolo/lite";

import "./static/index.scss";

import {home} from "./pages/home";

const app = okwolo();

app.setState({});

app("/", () => () => [home]);

app(/^\/(\w{6})$/g, (params) => () => (
    params[0]
));

app("**", () => () => (
    "404"
));

export default app;

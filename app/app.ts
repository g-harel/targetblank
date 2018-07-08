import okwolo from "okwolo/lite";

import "./static/index.scss";

import {home} from "./pages/home";
import {homepage} from "./pages/page";

const app = okwolo();

app.setState({});

app("/", () => () => (
    [home]
));

app(/^\/(\w{6})(?:\/(\S+?))?\/?$/g, (params) => () => (
    [homepage, {
        addr: params[0],
        token: params[1],
    }]
));

app("**", () => () => (
    "404"
));

export default app;

import okwolo from "okwolo/lite";

import "./static/index.scss";

import {landing} from "./pages/landing";
import {homepage, IHomepageProps} from "./pages/homepage";
import {password, IPasswordProps} from "./pages/password";

const app = okwolo();

app.setState({});

app("/", () => () => (
    [landing]
));

app(/^\/(\w{6})\/?$/g, (params) => () => (
    [homepage, {
        addr: params[0],
    } as IHomepageProps]
));

app(/^\/(\w{6})\/(\S+?)?\/?$/g, (params) => () => (
    [password, {
        addr: params[0],
        token: params[1],
    } as IPasswordProps]
));

app("**", () => () => (
    "404"
));

export default app;

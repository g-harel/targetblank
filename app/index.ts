import "./static/index.scss";

import {app} from "./app";
import {landing} from "./pages/landing";
import {homepage, IHomepageProps} from "./pages/homepage";
import {reset, IResetProps} from "./pages/reset";
import {login, ILoginProps} from "./pages/login";

app.use("target", document.body);

app.setState({});

app("/", () => () => (
    [landing]
));

app(/^\/(\w{6})\/?$/g, (params) => () => (
    [homepage, {
        addr: params[0],
    } as IHomepageProps]
));

app(/^\/(\w{6})\/login\/?$/g, (params) => () => (
    [login, {
        addr: params[0],
    } as ILoginProps]
));

app(/^\/(\w{6})\/reset(?:\/(\S+))?\/?/g, (params) => () => (
    [reset, {
        addr: params[0],
        token: params[1],
    } as IResetProps]
));

app("**", () => () => (
    "404"
));

import "./static/index.scss";

import {app} from "./app";
import {landing} from "./pages/landing";
import {homepage, IHomepageProps} from "./pages/homepage";
import {password, IPasswordProps} from "./pages/password";

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

app(/^\/(\w{6})\/(\S+?)?\/?$/g, (params) => () => (
    [password, {
        addr: params[0],
        token: params[1],
    } as IPasswordProps]
));

app("**", () => () => (
    "404"
));

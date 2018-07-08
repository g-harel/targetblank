import okwolo from "okwolo/lite";

import "./static/index.scss";

import {home} from "./pages/home";
import {homepage, IProps as IHomepageProps} from "./pages/homepage";

const app = okwolo();

app.setState({});

app("/", () => () => (
    [home]
));

app(/^\/(\w{6})(?:\/(\S+?))?\/?$/g, (params) => () => (
    [homepage, <IHomepageProps>{
        addr: params[0],
        token: params[1],
    }]
));

app("**", () => () => (
    "404"
));

export default app;

// The h function provided by okwolo is attached to the global object.
import h from "okwolo/src/h";
(window as any).h = h;

import "./global.scss";
import "normalize.css";

import {app} from "./internal/app";
import {routes} from "./routes";
import {Page} from "./components/page";

app.use("target", document.body);

app.setState({});

Object.keys(routes).forEach((name) => {
    const {path, component} = routes[name];
    app(path, (params) => () => <Page {...params} component={component} />);
});

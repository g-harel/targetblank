// The h function provided by okwolo is attached to the global object.
import h from "okwolo/src/h";
(window as any).h = h;

import "./global.css";
import "normalize.css";

import {app} from "./internal/app";
import {routes} from "./routes";
import {Page, Props} from "./components/page";

app.use("target", document.body);

app.setState({});

Object.keys(routes).forEach((name) => {
    const route = routes[name as keyof typeof routes];
    app(route.path, (params: Props) => () => <Page {...params} {...route} />);
});

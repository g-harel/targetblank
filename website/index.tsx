// The h function provided by okwolo is attached to the global object.
import h from "okwolo/src/h";
(window as any).h = h;

import "normalize.css";
import "./internal/style";

import {app} from "./internal/app";
import {registerRoutes} from "./routes";

app.use("target", document.body);

app.setState({});

registerRoutes();

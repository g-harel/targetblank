// The h function provided by okwolo is attached to the global object.
import h from "okwolo/src/h";
(window as any).h = h;

import "normalize.css";
import "./global.scss";

import {app} from "./internal/app";
import {Page} from "./components/page";
import {Landing} from "./pages/landing";
import {Document} from "./pages/document";
import {Forgot} from "./pages/forgot";
import {Edit} from "./pages/edit";
import {Reset} from "./pages/reset";
import {Login} from "./pages/login";
import {Missing} from "./pages/missing";

app.use("target", document.body);

app.setState({});

app("/", (params) => () => <Page {...params} component={Landing} />);

app("/:addr", (params) => () => <Page {...params} component={Document} />);

app("/:addr/edit", (params) => () => <Page {...params} component={Edit} />);

app("/:addr/login", (params) => () => <Page {...params} component={Login} />);

app("/:addr/forgot", (params) => () => <Page {...params} component={Forgot} />);

app("/:addr/reset/:token?", (params) => () => (
    <Page {...params} component={Reset} />
));

app("**", () => () => <Page component={Missing} />);

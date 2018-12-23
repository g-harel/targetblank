import "./static/index.scss";

import {app} from "./app";
import {Page} from "./components/page";
import {Landing} from "./pages/landing";
import {Homepage} from "./pages/homepage";
import {Reset} from "./pages/reset";
import {Login} from "./pages/login";
import {Missing} from "./pages/missing";

app.use("target", document.body);

app.setState({});

app("/", (params) => () => (
    <Page {...params} component={Landing} />
));

app("/:addr", (params) => () => (
    <Page {...params} component={Homepage} />
));

app("/:addr/login", (params) => () => (
    <Page {...params} component={Login} />
));

app("/:addr/reset/:token", (params) => () => (
    <Page {...params} component={Login} />
));

app("**", () => () => <Missing />);
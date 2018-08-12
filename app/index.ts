import "./static/index.scss";

import {app} from "./app";
import {page} from "./client/page";
import {landing} from "./pages/landing";
import {homepage, IHomepageProps} from "./pages/homepage";
import {reset, IResetProps} from "./pages/reset";
import {login, ILoginProps} from "./pages/login";

app.use("target", document.body);

app.setState({});

page("")(landing);

page<IHomepageProps>("{{addr}}", "addr")(homepage);

page<ILoginProps>("{{addr}}/login", "addr")(login);

page<IResetProps>("{{addr}}/reset{{token?}}", "addr", "token")(reset);

page("{{**}}")(() => () => "404");

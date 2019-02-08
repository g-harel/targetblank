import {app} from "./internal/app";
import {Component} from "./internal/types";
import {Landing} from "./pages/landing";
import {Document} from "./pages/document";
import {Recover} from "./pages/recover";
import {Edit} from "./pages/edit";
import {Reset} from "./pages/reset";
import {Login} from "./pages/login";
import {Missing} from "./pages/missing";

interface Route {
    path: string;
    component: Component<any, any>;
    allowLocalAddr?: boolean;
}

// Generates the route path from the path params.
export const path = (route: Route, ...params: string[]) => {
    let path = route.path;
    params.forEach((param) => {
        path = path.replace(/:\w+/, param);
    });
    path = path.replace(/:\w+\?$/, "");
    if (path.match(/:\w+/)) {
        throw "Missing redirect params";
    }
    return path;
};

// Redirects to the route using the path params.
export const redirect = (route: Route, ...params: string[]) => {
    app.redirect(path(route, ...params));
};

// Using an index signature ({[name: string]: Route}) would not keep hints about existing routes.
const routeTable = <T extends Record<string, Route>>(table: T): T => table;

export const routes = routeTable({
    landing: {
        path: "/",
        component: Landing,
    },
    document: {
        path: "/:addr",
        component: Document,
        allowLocalAddr: true,
    },
    edit: {
        path: "/:addr/edit",
        component: Edit,
        allowLocalAddr: true,
    },
    login: {
        path: "/:addr/login",
        component: Login,
    },
    recover: {
        path: "/:addr/recover",
        component: Recover,
    },
    reset: {
        path: "/:addr/reset/:token?",
        component: Reset,
    },
    missing: {
        path: "**",
        component: Missing,
    },
});

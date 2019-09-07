import {app} from "./internal/app";
import {Component} from "./internal/types";
import {Landing} from "./pages/landing";
import {Document} from "./pages/document";
import {Recover} from "./pages/recover";
import {Edit} from "./pages/edit";
import {Reset} from "./pages/reset";
import {Login} from "./pages/login";
import {Missing} from "./pages/missing";
import {Page, Props} from "./components/page";

const hostname = "targetblank.org";
const isExtension =
    (window as any).chrome &&
    (window as any).chrome.runtime &&
    (window as any).chrome.runtime.id;

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

// Unsafe relative redirect which works in an extension.
export const relativeRedirect = (path: string) => {
    if (isExtension) {
        const url = `https://${hostname}${path}`;
        (window as any).location = url;
    } else {
        app.redirect(path);
    }
};

// Redirects to the typed route using the path params.
export const safeRedirect = (route: Route, ...params: string[]) => {
    if (isExtension) {
        // TODO inject addr
        app.show(path(route, ...params));
    } else {
        app.redirect(path(route, ...params));
    }
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

export const registerRoutes = () => {
    Object.keys(routes).forEach((name) => {
        const route = routes[name as keyof typeof routes];
        app(route.path, (params: Props) => {
            if (isExtension) {
                // TODO infer from storage.
                params.addr = "test";
            }
            return () => <Page {...params} {...route} />;
        });
    });
};

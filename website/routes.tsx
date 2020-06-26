import {app} from "./internal/app";
import {Component} from "./internal/types";
import {Landing} from "./pages/landing";
import {Document} from "./pages/document";
import {Recover} from "./pages/recover";
import {Edit} from "./pages/edit";
import {Reset} from "./pages/reset";
import {Login} from "./pages/login";
import {Missing} from "./pages/missing";
import {Page, Props as PageProps} from "./components/page";
import {isExtension, read} from "./internal/extension";
import {Options} from "./pages/options";
import {Loading} from "./components/loading";
import {showChip} from "./components/page/chips";

const hostname = "targetblank.org";

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
export const relativeRedirect = async (path: string) => {
    if (isExtension) {
        const {addr} = await read();

        const canBeShown = !!(
            path === routes.options.path ||
            (addr && path.startsWith(`/${addr}`))
        );
        if (canBeShown) {
            app.show(path);
            return;
        }

        const url = `https://${hostname}${path}`;
        (window as any).location = url;
    } else {
        app.redirect(path);
    }
};

// Redirects to the typed route using the path params.
export const safeRedirect = (route: Route, ...params: string[]) => {
    if (isExtension) {
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
    options: {
        path: "/extension-options",
        component: Options,
    },
    extension: {
        // Matches the `window.location.pathname` for the extension's `newtab` page.
        // Must be placed before `/:addr` to have higher match priority.
        path: "/index.html",
        component: Document,
        allowLocalAddr: true,
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

// Transparently injects address when running in an extension context.
const PageLoader: Component<PageProps> = (params, update) => {
    let addr = params.addr;
    let addrIsLoading =
        isExtension &&
        addr === undefined &&
        params.component !== routes.options.component;

    if (addrIsLoading) {
        read().then(({addr: storedAddr}) => {
            if (storedAddr == null) {
                showChip("Please select your page", 6000);
                safeRedirect(routes.options);
                return;
            }
            addr = storedAddr;
            addrIsLoading = false;
            update();
        });
    }

    return () => {
        params.addr = addr;
        return addrIsLoading ? <Loading /> : <Page {...params} />;
    };
};

export const registerRoutes = () => {
    Object.keys(routes).forEach((name) => {
        const route = routes[name as keyof typeof routes];
        app(route.path, (params: PageProps) => {
            return () => <PageLoader {...params} {...route} />;
        });
    });
};

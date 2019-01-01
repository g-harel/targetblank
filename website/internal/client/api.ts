import {IPageData} from "./types";

const hostname = "https://api.targetblank.org";

interface IRequest {
    method: string;
    path: string;
    headers?: Record<string, string>;
    body?: any;
    json?: boolean;
}

// Helper to send the request using browser's fetch api.
// Also conditionally translates to and from json and appends hostname.
const send = async (req: IRequest) => {
    req.headers = req.headers || {};
    req.body = req.body || "";
    req.json = req.json || false;

    if (req.json) {
        req.body = JSON.stringify(req.body);
        req.headers["Content-Type"] = "application/json";
    }

    const res = await fetch(hostname + req.path, {
        method: req.method,
        headers: req.headers,
        body: req.body,
    });

    return req.json ? res.json() : res.text();
};

export const api = {
    page: {
        create: async (email: string): Promise<string> =>
            send({
                method: "POST",
                path: `/page`,
                body: email,
            }),
        delete: async (addr: string, token: string): Promise<void> =>
            send({
                method: "DELETE",
                path: `/page/${addr}`,
                headers: {token},
            }),
        edit: async (addr: string, token: string, spec: string): Promise<IPageData> =>
            send({
                method: "PUT",
                path: `/page/${addr}`,
                headers: {token},
                body: spec,
                json: true,
            }),
        fetch: async (addr: string, token: string): Promise<IPageData> =>
            send({
                method: "GET",
                path: `/page/${addr}`,
                headers: {token},
                json: true,
            }),
        publish: async (addr: string, token: string): Promise<void> =>
            send({
                method: "PATCH",
                path: `/page/${addr}`,
                headers: {token},
            }),
        validate: async (spec: string): Promise<void> =>
            send({
                method: "POST",
                path: `/page/validate`,
                body: spec,
            }),
        password: {
            change: async (addr: string, token: string, pass: string): Promise<void> =>
                send({
                    method: "PUT",
                    path: `/auth/${addr}`,
                    headers: {token},
                    body: pass,
                }),
            reset: async (addr: string, email: string): Promise<void> =>
                send({
                    method: "DELETE",
                    path: `/auth/${addr}`,
                    body: email,
                }),
        },
        token: {
            create: async (addr: string, pass: string): Promise<string> =>
                send({
                    method: "POST",
                    path: `/auth/${addr}`,
                    body: pass,
                }),
        },
    },
};

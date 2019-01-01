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
// Conditionally translates to and from json.
// Appends api's hostname.
// Rejects after a timeout period.
const send = async (req: IRequest) => {
    req.headers = req.headers || {};
    req.body = req.body || "";
    req.json = req.json || false;

    if (req.json) {
        req.body = JSON.stringify(req.body);
        req.headers["Content-Type"] = "application/json";
    }

    const res = await Promise.race([
        new Promise<never>((_, reject) => {
            setTimeout(() => {
                reject("Timeout");
            }, 5 * 1000);
        }),
        await fetch(hostname + req.path, {
            method: req.method,
            headers: req.headers,
            body: req.body,
        }),
    ]);

    return req.json ? res.json() : res.text();
};

export class PageAPI {
    async create(email: string): Promise<string> {
        return send({
            method: "POST",
            path: "/page",
            body: email,
        });
    }

    async delete(addr: string, token: string): Promise<void> {
        return send({
            method: "DELETE",
            path: `/page/${addr}`,
            headers: {token},
        });
    }

    async edit(addr: string, token: string, spec: string): Promise<IPageData> {
        return send({
            method: "PUT",
            path: `/page/${addr}`,
            headers: {token},
            body: spec,
            json: true,
        });
    }

    async fetch(addr: string, token: string): Promise<IPageData> {
        return send({
            method: "GET",
            path: `/page/${addr}`,
            headers: {token},
            json: true,
        });
    }

    async publish(addr: string, token: string): Promise<void> {
        return send({
            method: "PATCH",
            path: `/page/${addr}`,
            headers: {token},
        });
    }

    async validate(spec: string): Promise<void> {
        return send({
            method: "POST",
            path: "/page/validate",
            body: spec,
        });
    }
}

export class PagePasswordAPI {
    async change(addr: string, token: string, pass: string): Promise<void> {
        return send({
            method: "PUT",
            path: `/auth/${addr}`,
            headers: {token},
            body: pass,
        });
    }

    async reset(addr: string, email: string): Promise<void> {
        return send({
            method: "DELETE",
            path: `/auth/${addr}`,
            body: email,
        });
    }
}

export class PageTokenAPI {
    async create(addr: string, pass: string): Promise<string> {
        return send({
            method: "POST",
            path: `/auth/${addr}`,
            body: pass,
        });
    }
}

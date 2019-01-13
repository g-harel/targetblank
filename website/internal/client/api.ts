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
// Appends the correct host to the request path.
// Rejects after a timeout period.
// Rejects if the status code is in the error range.
const send = (req: IRequest) => {
    return new Promise<any>(async (resolve, reject) => {
        req.headers = req.headers || {};
        req.body = req.body || "";
        req.json = req.json || false;

        if (req.method === "GET" || req.method === "HEAD") {
            req.body = undefined;
        }

        if (req.json) {
            req.body = JSON.stringify(req.body);
            req.headers["Content-Type"] = "application/json";
        }

        // Time out request after interval.
        // All other resolve/rejects will have no effect.
        setTimeout(() => reject("Timed out"), 5 * 1000);

        const res = await fetch(hostname + req.path, {
            method: req.method,
            headers: req.headers,
            body: req.body,
        });

        // Message is forwarded to user-space only when status code is 400.
        if (res.status === 400) {
            return res.text().then(reject);
        }

        // Any other status code in the error range will have been encrypted.
        if (res.status > 400) {
            return res.text().then((message) => {
                console.error(`Status Code Error\n${message}`);
                reject("Something went wrong");
            });
        }

        if (req.json) {
            res.json().then(resolve);
        } else {
            res.text().then(resolve);
        }
    });
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

    async edit(addr: string, token: string, doc: string): Promise<IPageData> {
        return send({
            method: "PUT",
            path: `/page/${addr}`,
            headers: {token},
            body: doc,
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

    async validate(doc: string): Promise<void> {
        return send({
            method: "POST",
            path: "/page/validate",
            body: doc,
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

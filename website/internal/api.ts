import {IPageData} from "./types";

const hostname = "https://api.targetblank.org";

interface IRequest {
    method: string;
    path: string;
    token?: string;
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
        req.body = req.body || "";
        req.json = req.json || false;

        if (req.method === "GET" || req.method === "HEAD") {
            req.body = undefined;
        }

        // Arbitrary headers are not allowed because they would interfere with CORS.
        const headers = {};
        headers["Authorization"] = `Targetblank ${req.token}`;

        // Time out request after interval.
        // All other resolve/rejects will have no effect.
        setTimeout(() => reject("Timed out"), 5 * 1000);

        const res = await fetch(hostname + req.path, {
            headers,
            method: req.method,
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
    }).catch((e) => {
        console.warn(e.toString());
        throw e;
    });
};

export const pageCreate = async (email: string): Promise<string> => {
    return send({
        method: "POST",
        path: "/page",
        body: email,
    });
};

export const pageDelete = async (
    addr: string,
    token: string,
): Promise<void> => {
    return send({
        token,
        method: "DELETE",
        path: `/page/${addr}`,
    });
};

export const pageUpdate = async (
    addr: string,
    token: string,
    doc: string,
): Promise<IPageData> => {
    return send({
        token,
        method: "PUT",
        path: `/page/${addr}`,
        body: doc,
        json: true,
    });
};

export const pageRead = async (
    addr: string,
    token: string,
): Promise<IPageData> => {
    return send({
        token,
        method: "GET",
        path: `/page/${addr}`,
        json: true,
    });
};

export const pageValidate = async (doc: string): Promise<void> => {
    return send({
        method: "POST",
        path: "/page/validate",
        body: doc,
    });
};

export const tokenCreate = async (
    addr: string,
    pass: string,
): Promise<string> => {
    return send({
        method: "POST",
        path: `/auth/${addr}`,
        body: pass,
    });
};

export const passwordUpdate = async (
    addr: string,
    token: string,
    pass: string,
): Promise<void> => {
    return send({
        token,
        method: "PUT",
        path: `/auth/${addr}`,
        body: pass,
    });
};

export const passwordReset = async (
    addr: string,
    email: string,
): Promise<void> => {
    return send({
        method: "DELETE",
        path: `/auth/${addr}`,
        body: email,
    });
};

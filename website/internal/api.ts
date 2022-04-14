import {IPageData} from "./types";

const hostname = "https://api.targetblank.org";

interface IRequest {
    method: string;
    path: string;
    token?: string;
    body?: any;
    json?: boolean;
}

export interface IRequestError {
    isAuth?: boolean;
    isTimeout?: boolean;
    message: string;
}

// Helper to send the request using browser's fetch api.
// Conditionally translates to and from json.
// Appends the correct host to the request path.
// Rejects after a timeout period.
// Rejects if the status code is in the error range.
const send = (req: IRequest) => {
    return new Promise<any>(async (resolve, reject) => {
        const typedReject = (err: IRequestError) => reject(err);

        try {
            req.body = req.body || "";
            req.json = req.json || false;

            if (req.method === "GET" || req.method === "HEAD") {
                req.body = undefined;
            }

            // Arbitrary headers are not allowed because they would interfere with CORS.
            const headers: Record<string, string> = {};
            headers["Authorization"] = `Targetblank ${req.token}`;

            // Time out request after interval.
            // All other resolve/rejects will have no effect.
            setTimeout(() => {
                typedReject({
                    message: "request timeout",
                    isTimeout: true,
                });
            }, 5 * 1000);

            const res = await fetch(hostname + req.path, {
                headers,
                method: req.method,
                body: req.body,
            });

            // Likely auth related in the 4xx range.
            if (res.status >= 400 && res.status < 500) {
                return res.text().then((message) => {
                    typedReject({message, isAuth: true});
                });
            }

            // Unexpected errors.
            if (res.status >= 300) {
                return res.text().then((message) => {
                    console.error(`Status Code Error\n${message}`);
                    typedReject({message: "something went wrong"});
                });
            }

            if (req.json) {
                res.json().then(resolve);
            } else {
                res.text().then(resolve);
            }
        } catch (e) {
            // Fallback when function fails unexpectedly.
            console.warn(String(e));
            typedReject({message: String(e)});
        }
    });
};

export const pageCreate = async (email: string): Promise<string> => {
    return send({
        method: "POST",
        path: "/page",
        body: email,
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

export const pageValidate = async (doc: string): Promise<IPageData> => {
    return send({
        method: "POST",
        path: "/page/validate",
        body: doc,
        json: true,
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

export const passUpdate = async (
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

export const passReset = async (addr: string, email: string): Promise<void> => {
    return send({
        method: "DELETE",
        path: `/auth/${addr}`,
        body: email,
    });
};

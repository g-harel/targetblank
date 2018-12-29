import requestPromiseNative from "request-promise-native";

const endpoint = "https://api.targetblank.org";

export interface IPageItem {
    label: string;
    link: string;
    items: IPageItem[];
}

export interface IPageGroup {
    meta: {
        [key: string]: string;
    };
    items: IPageItem[];
}

export interface IPageData {
    version: string;
    spec: string;
    meta: {
        [key: string]: string;
    };
    groups: IPageGroup[];
}

export interface IError {
    statusCode: number;
    message: string;
}

export interface IAPI {
    page: {
        create(email: string): Promise<string>;
        delete(addr: string, token: string): Promise<void>;
        edit(addr: string, token: string, spec: string): Promise<IPageData>;
        fetch(addr: string, token?: string): Promise<IPageData>;
        publish(addr: string, token: string): Promise<void>;
        validate(spec: string): Promise<void>;
        password: {
            change(addr: string, token: string, pass: string): Promise<void>;
            reset(addr: string, email: string): Promise<void>;
        };
        token: {
            create(addr: string, pass: string): Promise<string>;
        };
    };
}

export const api: IAPI = {
    page: {
        create: async (email) =>
            requestPromiseNative({
                method: "POST",
                uri: `${endpoint}/page`,
                body: email,
            }),
        delete: async (addr, token) =>
            requestPromiseNative({
                method: "DELETE",
                uri: `${endpoint}/page/${addr}`,
                headers: {token},
            }),
        edit: async (addr, token, spec) =>
            requestPromiseNative({
                method: "PUT",
                uri: `${endpoint}/page/${addr}`,
                headers: {token},
                body: spec,
                json: true,
            }),
        fetch: async (addr, token) =>
            requestPromiseNative({
                method: "GET",
                uri: `${endpoint}/page/${addr}`,
                headers: {token},
                json: true,
            }),
        publish: async (addr, token) =>
            requestPromiseNative({
                method: "PATCH",
                uri: `${endpoint}/page/${addr}`,
                headers: {token},
            }),
        validate: async (spec) =>
            requestPromiseNative({
                method: "POST",
                uri: `${endpoint}/page/validate`,
                body: spec,
            }),
        password: {
            change: async (addr, token, pass) =>
                requestPromiseNative({
                    method: "PUT",
                    uri: `${endpoint}/auth/${addr}`,
                    headers: {token},
                    body: pass,
                }),
            reset: async (addr, email) =>
                requestPromiseNative({
                    method: "DELETE",
                    uri: `${endpoint}/auth/${addr}`,
                    body: email,
                }),
        },
        token: {
            create: async (addr, pass) =>
                requestPromiseNative({
                    method: "POST",
                    uri: `${endpoint}/auth/${addr}`,
                    body: pass,
                }),
        },
    },
};

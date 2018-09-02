import {api, IPageData} from "./api";
import {read, save} from "./storage";
import {show} from "./error";

type callback<T> = (result: T) => void;

export interface IClient {
    page: {
        create(email: string, fn: callback<string>);
        delete(addr: string, fn: callback<undefined>);
        edit(addr: string, spec: string, fn: callback<IPageData>);
        fetch(addr: string, fn: callback<IPageData>);
        publish(addr: string, fn: callback<undefined>);
        validate(spec: string, fn: callback<undefined>);
        password: {
            change(addr: string, pass: string, fn: callback<undefined>);
            reset(addr: string, email: string, fn: callback<undefined>);
        };
        token: {
            create(addr: string, pass: string, fn: callback<string>);
        };
    };
}

const missingToken = (addr: string) => show(`Missing access token for address ${addr}`);

export const client: IClient = {
    page: {
        create: async (email, fn) => {
            try {
                const res = await api.page.create(email);
                fn(res);
            } catch (e) {
                show(e.toString());
            }
        },
        delete: async (addr, fn) => {
            const {token} = read(addr);
            if (token === null) {
                return missingToken(addr);
            }
            try {
                await api.page.delete(addr, token);
                fn(undefined);
            } catch (e) {
                show(e.toString());
            }
        },
        token: {
            create: async (addr, pass, fn) => {
                try {
                    const token = await api.page.token.create(addr, pass);
                    save(addr, {token});
                    fn(token);
                } catch (e) {
                    show(e.toString());
                }
            },
        }
    }
};

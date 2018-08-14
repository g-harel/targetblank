import {api, IPageData} from "./api";
import {read, save} from "./storage";

type callback<T> = (result: T) => void;

export interface IClient {
    page: {
        create(email: string, fn: callback<string>);
        delete(addr: string, fn: callback<void>);
        edit(addr: string, spec: string, fn: callback<IPageData>);
        fetch(addr: string, fn: callback<IPageData>);
        publish(addr: string, fn: callback<void>);
        validate(spec: string, fn: callback<void>);
        password: {
            change(addr: string, pass: string, fn: callback<void>);
            reset(addr: string, email: string, fn: callback<void>);
        };
        token: {
            create(addr: string, pass: string, fn: callback<string>);
        };
    };
}

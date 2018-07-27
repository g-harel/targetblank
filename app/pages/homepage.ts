import "../static/page.landing.scss";

import {api, IError} from "../api";
import {password, IPasswordProps} from "../components/password";
import {read, save} from "../storage";

export interface IHomepageProps {
    addr: string;
}

const wrap = (element: any) => (
    ["div.homepage", {}, [
        element,
    ]]
);

export const homepage = ({addr}, update) => {
    let error: IError | null = null;

    let {data, token} = read(addr);

    api.page.fetch(addr, token)
        .then((d) => {
            data = d;
            save(addr, {data});
            update();
        })
        .catch((e) => {
            error = e;
            console.log(e);
            update();
        });

    return () => {
        if (data === null && error === null) {
            return wrap("loading");
        }

        return wrap(
            ["pre", {}, [
                error ? "couldn't load" : null,
                data && JSON.stringify(data, null, 2),
            ]],
        );
    };
};

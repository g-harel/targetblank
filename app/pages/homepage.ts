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
            update();
        });

    const submitPass = async (pass: string) => {
        try {
            token = await api.page.token.create(addr, pass);
            save(addr, {token});
            update();
        } catch (e) {
            console.log(e);
            return "Something went wrong";
        }
    };

    return () => {
        if (error !== null) {
            if (error.statusCode % 400 < 100) {
                return wrap(
                    [password, {
                        callback: submitPass,
                        title: "Sign in",
                    } as IPasswordProps],
                );
            }
            return wrap("an error has occured");
        }

        if (data === null) {
            return wrap("loading");
        }

        return wrap(
            ["pre", {}, [
                JSON.stringify(data, null, 2),
            ]],
        );
    };
};

import "../static/page.page.scss";

import {api, IError} from "../api";
import {input, props} from "../components/input";
import {read, save} from "../storage";

export interface IProps {
    addr: string;
    token?: string;
}

const passwordInput = (
    title: string,
    callback: props["callback"],
) => (
    ["div.password-input", {} , [
        [input, {
            title,
            callback,
            type: "password",
            validator: /.{8,32}/ig,
            message: "Password is too short",
            placeholder: "password123",
            focus: true,
        }],
    ]]
);

const wrap = (element: any) => (
    ["div.page", {}, [
        element,
    ]]
);

const passChange = (addr: string, token: string) => () => {
    const callback = async (pass: string) => {
        try {
            await api.page.password.change(addr, token, pass);
            window.location.pathname = "/" + addr;
        } catch (e) {
            console.log(e);
            return "Something went wrong";
        }
    };

    return wrap(passwordInput("Set your password", callback));
};

export const homepage = ({addr, token: tkn}, update) => {
    if (tkn) {
        return passChange(addr, tkn);
    }

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
                return wrap(passwordInput("Sign in", submitPass));
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

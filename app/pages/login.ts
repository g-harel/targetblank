import "../static/page.login.scss";

import {api} from "../api";
import {app} from "../app";
import {save} from "../storage";
import {password, IPasswordProps} from "../components/password";

export interface ILoginProps {
    addr: string;
}

export const login = ({addr}: ILoginProps) => () => {
    const callback = async (pass: string) => {
        try {
            const token = await api.page.token.create(addr, pass);
            save(addr, {token});
            app.redirect("/" + addr)
        } catch (e) {
            console.log(e);
            return "Something went wrong";
        }
    };

    return (
        ["div.login", {}, [
            [password, {
                callback,
                title: "Sign in",
            } as IPasswordProps],
        ]]
    );
};

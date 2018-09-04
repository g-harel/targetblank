import "../static/page.login.scss";

import {api} from "../client/api";
import {app} from "../app";
import {write} from "../client/storage";
import {password, IPasswordProps} from "../components/password";

export interface ILoginProps {
    addr: string;
}

export const login = ({addr}: ILoginProps) => () => {
    const callback = async (pass: string) => {
        try {
            const token = await api.page.token.create(addr, pass);
            write(addr, {token});
            app.redirect("/" + addr);
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

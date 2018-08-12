import "../static/page.reset.scss";

import {api} from "../client/api";
import {app} from "../app";
import {read} from "../client/storage";
import {password, IPasswordProps} from "../components/password";

export interface IResetProps {
    addr: string;
    token?: string;
}

export const reset = ({addr, token}: IResetProps) => () => {
    if (!token) {
        token = read(addr).token;
    }

    if (!token) {
        app.redirect("/" + addr + "/login");
    }

    const callback = async (pass: string) => {
        try {
            await api.page.password.change(addr, token, pass);
            app.redirect("/" + addr);
        } catch (e) {
            console.log(e);
            return "Something went wrong";
        }
    };

    return (
        ["div.password", {}, [
            [password, {
                callback,
                title: "Set your password",
            } as IPasswordProps],
        ]]
    );
};

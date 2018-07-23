import "../static/page.password.scss";

import {api} from "../api";
import {password as passwordComponent, IPasswordProps as IP} from "../components/password";

export interface IPasswordProps {
    addr: string;
    token: string;
}

export const password = ({addr, token}: IPasswordProps) => () => {
    const callback = async (pass: string) => {
        try {
            await api.page.password.change(addr, token, pass);
            window.location.pathname = "/" + addr;
        } catch (e) {
            console.log(e);
            return "Something went wrong";
        }
    };

    return (
        ["div.password", {}, [
            [passwordComponent, {
                callback,
                title: "Set your password",
            } as IP],
        ]]
    );
};

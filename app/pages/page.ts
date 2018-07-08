import {api, IError, IPageData} from "../api";
import {input} from "../components/input";

export interface IProps {
    addr: string;
    token?: string;
}

export const homepage = ({addr, token}, update) => {
    let data: IPageData | null = null;
    let error: IError | null = null;

    const load = () => {
        api.page.fetch(addr, token)
            .then((d) => {
                data = d;
                update();
            })
            .catch((e) => {
                error = e;
                update();
            });
    };

    const changePass = async (pass: string) => {
        try {
            await api.page.password.change(addr, token, pass);
            window.location.pathname = "/" + addr;
        } catch (e) {
            console.log(e);
            return "Something went wrong";
        }
    };

    return () => {
        let content = null;

        if (token) {
            content = (
                ["div.set-password", {}, [
                    [input, {
                        title: "Set your password",
                        type: "password",
                        callback: changePass,
                        validator: /.{8,32}/ig,
                        message: "Password is too short",
                        placeholder: "password123",
                    }],
                ]]
            );
        } else if (error !== null) {
            if (error.statusCode % 400 < 100) {
                content = "you need to log in";
            } else {
                content = "an error has occured";
            }
        } else if (data === null) {
            load();
            content = "loading";
        }
        return (
            ["div.page", {}, [
                content,
            ]]
        );
    };
};

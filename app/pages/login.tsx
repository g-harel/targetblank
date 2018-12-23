import "../static/page.login.scss";

import {api} from "../client/api";
import {app} from "../app";
import {write} from "../client/storage";
import {Password} from "../components/input/password";
import {PageComponent} from "../components/page";

export const Login: PageComponent = ({addr}) => () => {
    const callback = async (pass: string) => {
        try {
            const token = await api.page.token.create(addr, pass);
            write(addr, {token});
            app.redirect(`/${addr}`);
        } catch (e) {
            console.log(e);
            return "Something went wrong";
        }
    };

    return (
        <div className="login">
            <Password
                callback={callback}
                title="Sign in"
            />
        </div>
    );
};

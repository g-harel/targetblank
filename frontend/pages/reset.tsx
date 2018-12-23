import "../static/page.reset.scss";

import {api} from "../client/api";
import {app} from "../app";
import {read} from "../client/storage";
import {Password} from "../components/input/password";
import {PageComponent} from "../components/page";

export const Reset: PageComponent = ({addr, token}) => () => {
    if (!token) {
        token = read(addr).token;
    }

    if (!token) {
        app.redirect(`/${addr}/login`);
    }

    const callback = async (pass: string) => {
        try {
            await api.page.password.change(addr, token, pass);
            app.redirect(`/${addr}`);
        } catch (e) {
            console.log(e);
            return "Something went wrong";
        }
    };

    return (
        <div className="password">
            <Password title="Set your password" callback={callback} />
        </div>
    );
};
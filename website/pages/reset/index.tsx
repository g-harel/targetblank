import {api} from "../../internal/client/api";
import {app} from "../../internal/app";
import {read} from "../../internal/client/storage";
import {Password} from "../../components/input/password";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";

const Wrapper = styled("div")({
    paddingTop: "20vh",
});

export const Reset: PageComponent = ({addr, token: t}) => () => {
    const token = t || read(addr).token;

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
        <Wrapper>
            <Password title="Set your password" callback={callback} />
        </Wrapper>
    );
};

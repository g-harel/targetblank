import {api} from "../../library/client/api";
import {app} from "../../library/app";
import {write} from "../../library/client/storage";
import {Password} from "../../components/input/password";
import {PageComponent} from "../../components/page";
import {styled} from "../../library/styled";

const Wrapper = styled("div")({
    paddingTop: "20vh",
});

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
        <Wrapper>
            <Password
                callback={callback}
                title="Sign in"
            />
        </Wrapper>
    );
};

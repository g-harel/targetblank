import {client} from "../../internal/client";
import {app} from "../../internal/app";
import {Password} from "../../components/input/password";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";

const Wrapper = styled("div")({
    paddingTop: "20vh",
});

export const Login: PageComponent = ({addr}) => () => {
    const submit = (pass: string) => {
        return new Promise<string>((resolve) => {
            const callback = () => {
                resolve("");
                app.redirect(`/${addr}`);
            };

            const err = (message) => {
                resolve(message);
            };

            client.page.token.create(callback, err, addr, pass);
        });
    };

    return (
        <Wrapper>
            <Password callback={submit} title="Sign in" />
        </Wrapper>
    );
};

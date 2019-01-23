import {client} from "../../internal/client";
import {app} from "../../internal/app";
import {read, write} from "../../internal/client/storage";
import {Password} from "../../components/input/password";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";

const Wrapper = styled("div")({
    paddingTop: "20vh",
});

export const Reset: PageComponent = ({addr, token: t}) => () => {
    // Token from URL is used if it exists.
    const token = t || read(addr).token;
    write(addr, {token});

    const submit = (pass: string) => {
        return new Promise<string>((resolve) => {
            const callback = () => {
                resolve("");
                app.redirect(`/${addr}`);
            };

            const err = (msg) => {
                resolve(msg);
            };

            client.page.password.change(callback, err, addr, pass);
        });
    };

    return (
        <Wrapper>
            <Password title="Set your password" callback={submit} />
        </Wrapper>
    );
};

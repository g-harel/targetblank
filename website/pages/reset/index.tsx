import {client} from "../../internal/client";
import {app} from "../../internal/app";
import {Password} from "../../components/input/password";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";

const Wrapper = styled("div")({
    paddingTop: "20vh",
});

export const Reset: PageComponent = ({addr, token}) => () => {
    const submit = (pass: string) => {
        return new Promise<string>((resolve) => {
            client.page.password.change(
                () => app.redirect(`/${addr}/login`),
                resolve,
                addr,
                pass,
                token,
            );
        });
    };

    return (
        <Wrapper>
            <Password title="set password" callback={submit} />
        </Wrapper>
    );
};

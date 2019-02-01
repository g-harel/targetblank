import {client} from "../../internal/client";
import {app} from "../../internal/app";
import {Password} from "../../components/input/password";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";
import {Header} from "../../components/header";

const Wrapper = styled("div")({});

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
            <Header muted />
            <Password title="set password" validate callback={submit} />
        </Wrapper>
    );
};

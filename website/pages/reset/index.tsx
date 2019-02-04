import {client} from "../../internal/client";
import {Password} from "../../components/input/password";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";
import {Header} from "../../components/header";
import {routes, redirect} from "../../routes";

const Wrapper = styled("div")({});

export const Reset: PageComponent = ({addr, token}) => () => {
    const submit = (pass: string) => {
        return new Promise<string>((resolve) => {
            client.page.password.change(
                () => redirect(routes.login, addr),
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

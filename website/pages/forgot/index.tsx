import {client} from "../../internal/client";
import {Input} from "../../components/input";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/style";
import {Header} from "../../components/header";
import {routes, redirect} from "../../routes";

const Wrapper = styled("div")({});

export const Forgot: PageComponent = ({addr}) => () => {
    const submit = (email: string) => {
        return new Promise<string>((resolve) => {
            client(addr!).passReset(
                () => redirect(routes.login, addr!),
                resolve,
                email,
            );
        });
    };

    return (
        <Wrapper>
            <Header muted />
            <Input
                callback={submit}
                title="reset password"
                type="email"
                placeholder="john@example.com"
                validator={/.*/g}
                message=""
                focus={true}
            />
        </Wrapper>
    );
};

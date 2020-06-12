import {client} from "../../internal/client";
import {Password} from "../../components/input/password";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/style";
import {Header} from "../../components/header";
import {routes, safeRedirect} from "../../routes";

const Wrapper = styled("div")({});

export const Reset: PageComponent = ({addr, token}) => () => {
    // Token must be taken from URL to reset password.
    if (!token) {
        setTimeout(() => safeRedirect(routes.document, addr!));
        return null;
    }

    const submit = (pass: string) => {
        return new Promise<string>((resolve) => {
            client(addr!).passUpdate(
                () => {
                    client(addr!).tokenCreate(
                        () => safeRedirect(routes.document, addr!),
                        resolve,
                        pass,
                    );
                },
                resolve,
                pass,
                token,
            );
        });
    };

    return (
        <Wrapper>
            <Header muted />
            <Password
                title="set password"
                hint={addr}
                autocomplete="new-password"
                validate
                callback={submit}
            />
        </Wrapper>
    );
};

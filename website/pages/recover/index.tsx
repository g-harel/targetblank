import {client} from "../../internal/client";
import {Input} from "../../components/input";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/style";
import {Header} from "../../components/header";
import {routes, safeRedirect} from "../../routes";
import {handleErr} from "../../internal/errors";

const Wrapper = styled("div")({});

export const Recover: PageComponent = ({addr}) => () => {
    const submit = (email: string) => {
        return new Promise<string>((resolve) => {
            client(addr!).passReset(
                () => safeRedirect(routes.login, addr!),
                () => handleErr(resolve),
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
                hint={addr}
                type="email"
                autocomplete="email"
                placeholder="email@example.com"
                validator={/.*/g}
                message=""
                focus={true}
            />
        </Wrapper>
    );
};

import {client} from "../../internal/client";
import {Password} from "../../components/input/password";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/style";
import {color, size} from "../../internal/style/theme";
import {Anchor} from "../../components/anchor";
import {Header} from "../../components/header";
import {path, routes, safeRedirect} from "../../routes";

const Wrapper = styled("div")({});

const Recover = styled("div")({
    color: color.textSecondarySmall,
    fontSize: size.tiny,
    margin: "0 auto",
    paddingRight: "0.5rem",
    textAlign: "right",
    transform: "translateY(-1.85rem)",
    width: "16rem",
});

export const Login: PageComponent = ({addr}) => () => {
    const submit = (pass: string) => {
        return new Promise<string>((resolve) => {
            client(addr!).tokenCreate(
                () => safeRedirect(routes.document, addr!),
                resolve,
                pass,
            );
        });
    };

    return (
        <Wrapper>
            <Header muted />
            <Password
                callback={submit}
                title="log in"
                hint={addr}
                autocomplete="current-password"
            />
            <Recover>
                <Anchor id="reset" href={path(routes.recover, addr!)}>
                    reset password
                </Anchor>
            </Recover>
        </Wrapper>
    );
};

import {client} from "../../internal/client";
import {Password} from "../../components/input/password";
import {PageComponent} from "../../components/page";
import {styled, colors, size} from "../../internal/style";
import {Anchor} from "../../components/anchor";
import {Header} from "../../components/header";
import {path, routes, safeRedirect} from "../../routes";

const Wrapper = styled("div")({});

const Recover = styled("div")({
    color: colors.textSecondarySmall,
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
            <Password callback={submit} title="log in" />
            <Recover>
                <Anchor href={path(routes.recover, addr!)}>
                    reset password
                </Anchor>
            </Recover>
        </Wrapper>
    );
};

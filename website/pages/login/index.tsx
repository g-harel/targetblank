import {client} from "../../internal/client";
import {Password} from "../../components/input/password";
import {PageComponent} from "../../components/page";
import {styled, colors} from "../../internal/style";
import {Anchor} from "../../components/anchor";
import {Header} from "../../components/header";
import {path, routes, redirect} from "../../routes";

const Wrapper = styled("div")({});

const Recover = styled("div")({
    color: colors.textSecondarySmall,
    fontSize: "0.8rem",
    margin: "0 auto",
    paddingRight: "0.5rem",
    textAlign: "right",
    transform: "translateY(-2rem)",
    width: "16rem",
});

export const Login: PageComponent = ({addr}) => () => {
    const submit = (pass: string) => {
        return new Promise<string>((resolve) => {
            client(addr!).tokenCreate(
                () => redirect(routes.document, addr!),
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

import {client} from "../../internal/client";
import {app} from "../../internal/app";
import {Password} from "../../components/input/password";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";
import {Anchor} from "../../components/anchor";
import {Header} from "../../components/header";

const Wrapper = styled("div")({});

const Forgot = styled("div")({
    color: "#aaa",
    fontSize: "0.8rem",
    margin: "0 auto",
    paddingRight: "0.5rem",
    textAlign: "right",
    transform: "translateY(-1.9rem)",
    width: "16rem",
});

export const Login: PageComponent = ({addr}) => () => {
    const submit = (pass: string) => {
        return new Promise<string>((resolve) => {
            client.page.token.create(
                () => app.redirect(`/${addr}`),
                resolve,
                addr,
                pass,
            );
        });
    };

    return (
        <Wrapper>
            <Header muted />
            <Password callback={submit} title="log in" />
            <Forgot>
                <Anchor href={`/${addr}/forgot`}>reset password</Anchor>
            </Forgot>
        </Wrapper>
    );
};

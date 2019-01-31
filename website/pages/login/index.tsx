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
    cursor: "pointer",
    fontSize: "0.9rem",
    height: 0,
    margin: "0 auto",
    padding: "0 0.9rem",
    textAlign: "right",
    transform: "translateY(0.85rem)",
    userSelect: "none",
    width: "17rem",
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
            <Forgot>
                <Anchor href={`/${addr}/forgot`}>reset password</Anchor>
            </Forgot>
            <Password callback={submit} title="log in" />
        </Wrapper>
    );
};

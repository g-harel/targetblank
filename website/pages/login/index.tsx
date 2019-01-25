import {client} from "../../internal/client";
import {app} from "../../internal/app";
import {Password} from "../../components/input/password";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";

const Wrapper = styled("div")({
    paddingTop: "20vh",
});

const ForgotWrapper = styled("div")({
    height: 0,
    margin: "0 auto",
    padding: "0 0.9rem",
    textAlign: "right",
    transform: "translateY(0.85rem)",
    width: "17rem",
});

const ForgotLink = styled("div")({
    color: "#aaa",
    cursor: "pointer",
    display: "inline-block",
    fontSize: "0.9rem",
    userSelect: "none",
});

export const Login: PageComponent = ({addr}) => () => {
    const submit = (pass: string) => {
        return new Promise<string>((resolve) => {
            client.page.token.create(
                () => app.redirect(`/${addr}/login`),
                resolve,
                addr,
                pass,
            );
        });
    };

    const onClick = () => app.redirect(`/${addr}/forgot`);

    return (
        <Wrapper>
            <ForgotWrapper>
                <ForgotLink onclick={onClick}>
                    reset password
                </ForgotLink>
            </ForgotWrapper>
            <Password callback={submit} title="log in" />
        </Wrapper>
    );
};

import {client} from "../../internal/client";
import {app} from "../../internal/app";
import {Input} from "../../components/input";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";
import {Header} from "../../components/header";

const Wrapper = styled("div")({});

export const Forgot: PageComponent = ({addr}) => () => {
    const submit = (email: string) => {
        return new Promise<string>((resolve) => {
            client.page.password.reset(
                () => app.redirect(`/${addr}`),
                resolve,
                addr,
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

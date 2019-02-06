import {client} from "../../internal/client";
import {Password} from "../../components/input/password";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/style";
import {Header} from "../../components/header";
import {routes, redirect} from "../../routes";

const Wrapper = styled("div")({});

export const Reset: PageComponent = ({addr, token}) => () => {
    const submit = (pass: string) => {
        return new Promise<string>((resolve) => {
            client(addr).passUpdate(
                () => {
                    client(addr).tokenCreate(
                        () => redirect(routes.document, addr),
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
            <Password title="set password" validate callback={submit} />
        </Wrapper>
    );
};

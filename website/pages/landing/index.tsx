import {client, localAddr} from "../../internal/client";
import {styled} from "../../internal/style";
import {Signup} from "./signup";
import {PageComponent} from "../../components/page";
import {Header} from "../../components/header";
import {Confirmation} from "./confirmation";
import {Anchor} from "../../components/anchor";
import {path, routes} from "../../routes";

const screenWidth = 20;

const Wrapper = styled("div")({
    overflow: "hidden",
});

const Screens = styled("div")({
    marginLeft: `calc(50vw - ${screenWidth / 2}rem);`,
    width: "1000vw",
});

const Screen = styled("div")({
    display: "inline-block",
    opacity: 0,
    textAlign: "center",
    transition: "all 0.7s ease",
    verticalAlign: "top",
    width: `${screenWidth}rem`,

    $nest: {
        "&.scrolled": {
            transform: "translateX(-100%)",
        },

        "&.visible": {
            opacity: 1,
        },
    },
});

const Info = styled("div")({
    fontSize: "0.8rem",
    fontWeight: 600,
    padding: "0 1rem",
    marginTop: "-0.5rem",
});

const Spacer = styled("div")({
    height: 0,
    borderBottom: "1px solid #eee",
    margin: "2rem",
});

const Try = styled("div")({
    color: "#bbb",
    fontWeight: 600,
    marginBottom: "1.9rem",
});

export const Landing: PageComponent = (_, update) => {
    let email = "";

    const submit = (newEmail: string) => {
        return new Promise<string>((resolve) => {
            const callback = () => {
                email = newEmail;
                update();
                resolve("");
            };
            client.pageCreate(callback, resolve, newEmail);
        });
    };

    return () => (
        <Wrapper>
            <Header sub="organize your links" />
            <Screens>
                <Screen className={{scrolled: !!email, visible: !email}}>
                    <Signup callback={submit} />
                    <Info>
                        Your email address is stored hashed, like a password.
                        You will be asked to re-enter it to reset the page's
                        password.
                    </Info>
                    <Spacer />
                    <Try>
                        <Anchor href={path(routes.document, localAddr)}>
                            try it out locally
                        </Anchor>
                    </Try>
                    <Info>
                        The local page can only be accessed from this computer.
                        Document must still be parsed on the server.
                    </Info>
                </Screen>
                <Screen className={{scrolled: !!email, visible: !!email}}>
                    <Confirmation email={email} />
                </Screen>
            </Screens>
        </Wrapper>
    );
};

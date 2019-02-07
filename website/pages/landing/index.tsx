import {client} from "../../internal/client";
import {styled} from "../../internal/style";
import {Signup} from "./signup";
import {PageComponent} from "../../components/page";
import {Header} from "../../components/header";
import {Confirmation} from "./confirmation";

const screenWidth = 20;

const Wrapper = styled("div")({
    overflowX: "hidden",
});

const Screens = styled("div")({
    marginLeft: `calc(50vw - ${screenWidth / 2}rem);`,
    width: "1000vw",
});

const Screen = styled("div")({
    display: "inline-block",
    textAlign: "center",
    transition: "all 0.7s ease",
    verticalAlign: "top",
    width: `${screenWidth}rem`,

    $nest: {
        "&.scrolled": {
            transform: "translateX(-100%)",
        },
    },
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
            <Header />
            <Screens>
                <Screen className={{scrolled: !!email}}>
                    <Signup callback={submit} visible={!email} />
                </Screen>
                <Screen className={{scrolled: !!email}}>
                    <Confirmation email={email} visible={!!email} />
                </Screen>
            </Screens>
        </Wrapper>
    );
};

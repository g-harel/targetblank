import {api} from "../../library/client/api";
import {styled} from "../../library/styled";
import {Signup} from "./signup";
import {PageComponent} from "../../components/page";
import {Footer, footerHeight} from "./footer";
import {Header, headerHeight} from "./header";
import {Confirmation} from "./confirmation";

const screenWidth = 20;

const Wrapper = styled("div")({
    overflowX: "hidden",
});

const Screens = styled("div")({
    height: `calc(100vh - ${headerHeight + footerHeight + 0.1}rem)`,
    marginLeft: `calc(50vw - ${screenWidth / 2}rem);`,
    width: "1000vw",
});

const Screen = styled("div")({
    display: "inline-block",
    height: "99%",
    paddingTop: `calc(25vh - ${headerHeight}rem)`,
    textAlign: "center",
    transition: "all 0.7s ease",
    verticalAlign: "top",
    width: `${screenWidth}rem`,

    "&.scrolled": {
        transform: "translateX(-100%)",
    },
});

export const Landing: PageComponent = (props, update) => {
    let email = "";

    const callback = async (addr: string) => {
        try {
            await api.page.create(addr);
        } catch (e) {
            console.log(e);
            return "Something went wrong";
        }

        email = addr;
        update();
    };

    return () => (
        <Wrapper>
            <Header />
            <Screens>
                <Screen className={{scrolled: !!email}}>
                    <Signup callback={callback} visible={!email} />
                </Screen>
                <Screen className={{scrolled: !!email}}>
                    <Confirmation email={email} visible={!!email} />
                </Screen>
            </Screens>
            <Footer />
        </Wrapper>
    );
};

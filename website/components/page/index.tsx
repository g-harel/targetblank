import {Missing} from "../../pages/missing";
import {Component} from "../../internal/types";
import {styled} from "../../internal/style";
import {color, size} from "../../internal/style/theme";
import {reset} from "../../internal/keyboard";
import {localAddr} from "../../internal/client";
import {Anchor} from "../anchor";
import {path, routes} from "../../routes";
import {isExtension} from "../../internal/extension";
import {Chips} from "./chips";

const Wrapper = styled("div")({
    display: "flex",
    flexDirection: "column",
    minHeight: "100%",
});

const Content = styled("main")({
    display: "flex",
    flex: "1 0 auto",
    flexDirection: "column",
});

const Footer = styled("footer")({
    color: color.textSecondaryLarge,
    flexShrink: 0,
    fontSize: size.tiny,
    fontWeight: 600,
    padding: "1.5rem",
    textAlign: "center",
});

const ExtensionOptions = styled("div")({
    padding: "1rem",
    color: color.textPrimary,
});

export interface PageProps {
    addr?: string;
    token?: string;
}

export type PageComponent<A = any> = Component<PageProps, A>;

export interface Props extends PageProps {
    component: PageComponent;
    allowLocalAddr?: boolean;
}

export const Page: Component<Props> = (props) => () => {
    const {addr, token, allowLocalAddr} = props;
    let Component: PageComponent = props.component;

    if (addr === localAddr) {
        if (!allowLocalAddr) {
            console.warn("local address not allowed");
            Component = Missing;
        }
    } else if (addr && !addr.match(/^\w{6}$/g)) {
        console.warn("invalid `addr` in path");
        Component = Missing;
    }
    if (token && !token.match(/^[^\s\/]+$/)) {
        console.warn("invalid `token` in path");
        Component = Missing;
    }

    // Reset keyboard listeners added by the previous page.
    reset();

    // Reset document title.
    document.title = "targetblank";

    return (
        <Wrapper>
            <Chips />
            <Content>
                <Component addr={addr} token={token} />
            </Content>
            <Footer>
                {isExtension && (
                    <ExtensionOptions>
                        <Anchor href={path(routes.options)}>
                            extension options
                        </Anchor>
                    </ExtensionOptions>
                )}
                <Anchor id="home" href="/">
                    home
                </Anchor>
                &nbsp;&nbsp;&nbsp;
                <Anchor
                    id="about"
                    href="https://github.com/g-harel/targetblank/#readme"
                >
                    about
                </Anchor>
                &nbsp;&nbsp;&nbsp;
                <Anchor id="contact" href="mailto:gabrielj.harel@gmail.com">
                    contact
                </Anchor>
            </Footer>
        </Wrapper>
    );
};

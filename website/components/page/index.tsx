import {Missing} from "../../pages/missing";
import {Component} from "../../internal/types";
import {styled, colors, size} from "../../internal/style";
import {reset} from "../../internal/keyboard";
import {localAddr} from "../../internal/client";
import {Anchor} from "../anchor";

const Wrapper = styled("div")({
    backgroundColor: colors.backgroundSecondary,
    display: "flex",
    flexDirection: "column",
    minHeight: "100%",
});

const Content = styled("div")({
    flex: "1 0 auto",
});

const Footer = styled("footer")({
    backgroundColor: colors.backgroundSecondary,
    color: colors.textSecondaryLarge,
    flexShrink: 0,
    fontSize: size.tiny,
    fontWeight: 600,
    padding: "1.5rem",
    textAlign: "center",
});

const ConstructionSign = styled("div")({
    backgroundColor: colors.error,
    borderRadius: "0.3rem",
    bottom: "2rem",
    color: colors.backgroundPrimary,
    fontFamily: "monospace",
    fontWeight: 600,
    left: "2rem",
    padding: "1rem",
    position: "fixed",
    textAlign: "center",
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
            {document.location.hostname !== "localhost" && (
                <ConstructionSign>Under Construction</ConstructionSign>
            )}
            <Content>
                <Component addr={addr} token={token} />
            </Content>
            <Footer>
                <Anchor href="/">home</Anchor>
                &nbsp;&nbsp;&nbsp;
                <Anchor href="https://github.com/g-harel/targetblank/#readme">
                    about
                </Anchor>
                &nbsp;&nbsp;&nbsp;
                <Anchor href="mailto:gabrielj.harel@gmail.com">contact</Anchor>
            </Footer>
        </Wrapper>
    );
};

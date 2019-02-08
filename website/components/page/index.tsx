import {Missing} from "../../pages/missing";
import {Component} from "../../internal/types";
import {styled} from "../../internal/style";
import {reset} from "../../internal/keyboard";
import {localAddr} from "../../internal/client";

const Wrapper = styled("div")({
    height: "100%",
});

const ConstructionSign = styled("div")({
    backgroundColor: "red",
    borderRadius: "0.3rem",
    bottom: "2rem",
    color: "white",
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
            console.warn("local `addr` not allowed");
            Component = Missing;
        }
    } else if (addr && !addr.match(/^\w{6}$/)) {
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
            <Component addr={addr} token={token} />
        </Wrapper>
    );
};

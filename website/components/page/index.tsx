import {Missing} from "../../pages/missing";
import {Component} from "../../internal/types";
import {styled} from "../../internal/style";
import {reset} from "../../internal/keyboard";

const Wrapper = styled("div")({
    height: "100%",
});

const ConstructionSign = styled("div")({
    backgroundColor: "red",
    borderRadius: "0.3rem",
    bottom: "2rem",
    color: "white",
    fontFamily: "monospace",
    fontWeight: "bold",
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
}

export const Page: Component<Props> = (props) => () => {
    let Component: PageComponent = props.component;

    if (props.addr && !props.addr.match(/^\w{6}$/)) {
        console.warn("invalid `addr` in path");
        Component = Missing;
    }
    if (props.token && !props.token.match(/^[^\s\/]+$/)) {
        console.warn("invalid `token` in path");
        Component = Missing;
    }

    // Reset keyboard listeners added by the previous page.
    reset();

    return (
        <Wrapper>
            {document.location.hostname !== "localhost" && (
                <ConstructionSign>Under Construction</ConstructionSign>
            )}
            <Component addr={props.addr} token={props.token} />
        </Wrapper>
    );
};

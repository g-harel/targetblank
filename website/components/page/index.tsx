import {ErrorComponent} from "../../internal/client/error";
import {Missing} from "../../pages/missing";
import {Component} from "../../internal/types";
import {styled} from "../../internal/styled";

const Wrapper = styled("div")({
    backgroundColor: "gold",
    borderRadius: "0.3rem",
    bottom: "2rem",
    fontFamily: "Inconsolata, monospace",
    fontSize: "1.2rem",
    fontWeight: "bold",
    left: "2rem",
    opacity: 0.5,
    padding: "1rem",
    position: "fixed",
    textAlign: "center",
});

export interface PageProps {
    addr?: string;
    token?: string;
}

export type PageComponent<A = undefined> = Component<PageProps, A>;

export interface Props extends PageProps {
    component: PageComponent;
}

export const Page: Component<Props> = (props) => () => {
    let Component: PageComponent = props.component;

    if (props.addr && !props.addr.match(/\w{6}/)) {
        console.warn("invalid `addr` in path");
        Component = Missing;
    }
    if (props.token && !props.token.match(/[^\s\/]+/)) {
        console.warn("invalid `token` in path");
        Component = Missing;
    }

    return (
        <div className="page">
            {document.location.hostname !== "localhost" && (
                <Wrapper>
                    <i className="fas fa-tools" />
                    &nbsp; Under Construction
                </Wrapper>
            )}
            <Component addr={props.addr} token={props.token} />
            <ErrorComponent />
        </div>
    );
};

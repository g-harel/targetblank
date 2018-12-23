import {ErrorComponent} from "../../library/client/error";
import {Missing} from "../../pages/missing";
import {Component} from "../../library/types";

export interface PageProps {
    addr?: string;
    token?: string;
}

export type PageComponent<A = undefined> = Component<PageProps, A>;

export interface Props extends PageProps {
    component: PageComponent;
}

export const Page: Component<Props> = (props) => () => {
    if (props.addr && !props.addr.match(/\w{6}/)) {
        console.warn("invalid `addr` in path");
        return <Missing />;
    }
    if (props.token && !props.token.match(/[^\s\/]+/)) {
        console.warn("invalid `token` in path");
        return <Missing />;
    }

    const Component: PageComponent = props.component;
    return (
        <div className="page">
            <Component
                addr={props.addr}
                token={props.token}
            />
            <ErrorComponent />
        </div>
    );
};

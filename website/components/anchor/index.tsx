import {Component} from "../../internal/types";
import {app} from "../../internal/app";
import {styled} from "../../internal/styled";

const Wrapper = styled("a")({
    color: "inherit",
});

export interface Props {
    href: string;
}

export const Anchor: Component<Props> = (props) => {
    const onClick = (e) => {
        if (props.href && (props.href[0] === "/" || props.href[0] === ".")) {
            e.preventDefault();
            app.redirect(props.href);
        }
    };

    return () => (
        <Wrapper href={props.href} onclick={onClick}>
            {...(props as any).children}
        </Wrapper>
    );
};

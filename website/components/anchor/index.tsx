import {Component} from "../../internal/types";
import {app} from "../../internal/app";
import {styled} from "../../internal/style";

const Wrapper = styled("a")({
    color: "inherit",
});

const isRelative = (href: String): boolean => {
    return href[0] === "/" || href[0] === ".";
};

export interface Props {
    href: string;
}

export const Anchor: Component<Props> = (props) => {
    let href = props.href || "";

    // Assume "https" protocol if none specified.
    if (!isRelative(href) && href.match(/\./) && !href.match(/^\w+:/)) {
        href = `https://${href}`;
    }

    const onClick = (e) => {
        if (isRelative(href)) {
            e.preventDefault();
            app.redirect(href);
        }
    };

    return () => (
        <Wrapper href={href} onclick={onClick}>
            {...(props as any).children}
        </Wrapper>
    );
};

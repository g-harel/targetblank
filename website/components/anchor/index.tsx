import {Component} from "../../internal/types";
import {styled} from "../../internal/style";
import {relativeRedirect} from "../../routes";

const Wrapper = styled("a")({
    color: "inherit",
});

const isRelative = (href: String): boolean => {
    return href[0] === "/" || href[0] === ".";
};

export interface Props {
    id?: string;
    href: string;
    immediate?: boolean;
}

export const Anchor: Component<Props> = (props) => {
    let href = props.href || "";

    // Assume "https" protocol if none specified.
    if (!isRelative(href) && href.match(/\./) && !href.match(/^\w+:/)) {
        href = `https://${href}`;
    }

    const onClick = (e: MouseEvent) => {
        if (isRelative(href)) {
            e.preventDefault();
            relativeRedirect(href);
        }
    };

    // Immediately follow link.
    if (props.immediate) {
        if (isRelative(href)) {
            relativeRedirect(href);
        } else {
            (window as any).location = href;
        }
    }

    return () => (
        <Wrapper id={props.id} href={href} onclick={onClick}>
            {...(props as any).children}
        </Wrapper>
    );
};

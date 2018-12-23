import {styled} from "../../library/styled";
import {Component} from "../../library/types";

// most common email providers and their web interface url
const providers: Record<string, string> = {
    gmail: "https://mail.google.com",
    icloud: "https://mail.icloud.com",
    yahoo: "https://mail.yahoo.com",
    outlook: "https://outlook.live.com",
};

// recognized email domains and associated provider
const domains: Record<string, string> = {
    gmail: "gmail",
    icloud: "icloud",
    me: "icloud",
    yahoo: "yahoo",
    live: "outlook",
    outlook: "outlook",
    hotmail: "outlook",
    msn: "outlook",
};

const Wrapper = styled("div")({
    opacity: 0,

    "&.visible": {
        opacity: 1,
    },
});

const Title = styled("span")({
    fontSize: "1.3rem",
    fontWeight: "bold",
});

const TitleIcon = styled("i")({
    color: "yellowgreen",
    display: "block",
    margin: "0.6rem 0",
    opacity: 0,
    transition: "all 1s ease",

    "&.visible": {
        opacity: 1,
        transitionDelay: "0.25s",
    },
});

const EmailLink = styled("a")({
    color: "#aaa",
    display: "inline-block",
    margin: "0.5rem 0 1rem",

    "&.inert": {
        cursor: "text",
        marginTop: "0.3rem",
        pointerEvents: "none",
    },
});

const EmailLinkIcon = styled("i")({
    fontSize: "0.6rem",
    opacity: 0.9,
    paddingLeft: "0.25rem",
    transform: "translateY(-0.5px)",
});

const Providers = styled("div")({
    paddingTop: "0.8rem",
});

const ProviderLink = styled("a")({
    borderLeft: "1px solid #eee",
    color: "#aaa",
    display: "inline-block",
    maxWidth: "25%",
    padding: "0 0.5rem",
    transition: "all 0.5s ease",

    "&:hover": {
        color: "#888",
    },

    "&:first-child": {
        border: "none",
    },
});

export interface Props {
    email: string;
    visible?: boolean;
}

export const Confirmation: Component<Props> = (props) => {
    const {email, visible} = props;

    // attempt to find the web interface link from the email's domain
    let link: string | null = null;
    const match = /.*@([^\.]+).*/g.exec(email);
    if (match !== null) {
        const [, domain] = match;
        link = providers[domains[domain]] || null;
    }

    return () => (
        <Wrapper className={{visible}}>
            <Title>
                <TitleIcon className={[
                    "far fa-lg fa-check-circle",
                    {visible},
                ]} />
                Confirmation Sent
            </Title>
            <div>
                <EmailLink href={link} className={{inert: !link}}>
                    {email}
                    {link && (
                        <EmailLinkIcon className="fas fa-xs fa-external-link-alt" />
                    )}
                </EmailLink>
                {!link && (
                    <Providers>
                        {...Object.keys(providers).map((name) => (
                            <ProviderLink href={providers[name]}>
                                {name}
                            </ProviderLink>
                        ))}
                    </Providers>
                )}
            </div>
        </Wrapper>
    );
};

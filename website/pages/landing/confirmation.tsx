import {styled} from "../../internal/styled";
import {Component} from "../../internal/types";

// Most common email providers and their web interface url.
const providers: Record<string, string> = {
    gmail: "https://mail.google.com",
    icloud: "https://mail.icloud.com",
    yahoo: "https://mail.yahoo.com",
    outlook: "https://outlook.live.com",
};

// Recognized email domains and associated provider.
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
    fontSize: "1.1rem",
    margin: "0.5rem 0 1rem",

    "&.inert": {
        cursor: "text",
        marginTop: "0.3rem",
        pointerEvents: "none",
    },
});

export interface Props {
    email: string;
    visible?: boolean;
}

export const Confirmation: Component<Props> = (props) => {
    const {email, visible} = props;

    // Attempt to find the web interface link from the email's domain.
    let link: string | null = null;
    const match = /.*@([^\.]+).*/g.exec(email);
    if (match !== null) {
        const [, domain] = match;
        link = providers[domains[domain]] || null;
    }

    return () => (
        <Wrapper className={{visible}}>
            <Title>
                <TitleIcon
                    className={["far fa-lg fa-check-circle", {visible}]}
                />
                Confirmation Sent
            </Title>
            <div>
                <EmailLink href={link} className={{inert: !link}}>
                    {email}
                </EmailLink>
            </div>
        </Wrapper>
    );
};

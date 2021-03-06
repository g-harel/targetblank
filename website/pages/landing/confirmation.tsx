import {styled} from "../../internal/style";
import {color, size} from "../../internal/style/theme";
import {Component} from "../../internal/types";
import {Icon} from "../../components/icon";

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

const Wrapper = styled("div")({});

const Title = styled("span")({
    fontSize: size.large,
    fontWeight: 600,
});

const TitleIcon = styled("div")({
    margin: "1rem 0 0.5rem",
    transition: "all 1s ease",
});

const EmailLink = styled("a")({
    color: color.textSecondarySmall,
    fontSize: size.tiny,
    display: "inline-block",
    margin: "0.5rem 0 1rem",

    $nest: {
        "&.inert": {
            cursor: "text",
            marginTop: "0.3rem",
            pointerEvents: "none",
        },
    },
});

export interface Props {
    email: string;
}

export const Confirmation: Component<Props> = (props) => {
    const {email} = props;

    // Attempt to find the web interface link from the email's domain.
    let link: string | null = null;
    const match = /.*@([^\.]+).*/g.exec(email);
    if (match !== null) {
        const [, domain] = match;
        link = providers[domains[domain]] || null;
    }

    return () => (
        <Wrapper>
            <Title>
                <TitleIcon>
                    <Icon name="check" color="yellowgreen" size={1.4} />
                </TitleIcon>
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

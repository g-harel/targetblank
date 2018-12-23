import "../static/page.landing.scss";

import {api} from "../client/api";
import {styled} from "../styled";
import {Input} from "../components/input";
import {PageComponent} from "../components/page";

type email = null | {
    addr: string;
    link: string | null;
};

const providers: {
    [p: string]: {
        domains: string[],
        url: string,
    },
} = {
    gmail: {
        domains: ["gmail"],
        url: "https://mail.google.com",
    },
    icloud: {
        domains: ["icloud"],
        url: "https://mail.icould.com",
    },
    yahoo: {
        domains: ["yahoo"],
        url: "https://mail.yahoo.com",
    },
    outlook: {
        domains: ["live", "outlook", "hotmail", "msn"],
        url: "https://outlook.live.com",
    },
};

const calcEmail = (addr: string): email => {
    const match = /.*@([^\.]+).*/g.exec(addr);
    if (match === null) {
        return {addr, link: null};
    }
    const [, domain] = match;

    let link = null;
    Object.keys(providers).forEach((p) => {
        const {domains, url} = providers[p];
        if (domains.indexOf(domain) >= 0) {
            link = url;
        }
    });

    return {addr, link};
};

const TitleHighlight = styled("span")({
    color: "grey" || "#ffbd00",
    font: "inherit",
});

export const Landing: PageComponent = (props, update) => {
    let email: email = null;
    let scrolled = false;

    const callback = async (addr: string) => {
        try {
            await api.page.create(addr);
        } catch (e) {
            console.log(e);
            return "Something went wrong";
        }

        scrolled = true;
        email = calcEmail(addr);
        update();
    };

    return () => (
        <div className="landing">
            <div className="header">
                target
                <TitleHighlight>
                    blank
                </TitleHighlight>
            </div>
            <div className={{scrolled, screens: true}}>
                <div className="screen signup">
                    <Input
                        callback={callback}
                        title="Create a homepage"
                        placeholder="john@example.com"
                        validator={/^\S+@\S+\.\S{2,}$/g}
                        message="That doesn't look like an email address"
                        focus={true}
                    />
                </div>
                <div className="screen confirmation">
                    <span className="title">
                        <i className="far fa-lg fa-check-circle" />
                        Confirmation Sent
                    </span>
                    {email && (
                        <div className="email">
                            <a href={email.link} className={{
                                address: true,
                                inert: !email.link,
                            }}>
                                {email.addr}
                                {email.link && (
                                    <i className="fas fa-xs fa-external-link-alt" />
                                )}
                            </a>
                            {!email.link && (
                                <div className="providers">
                                    {Object.keys(providers).map((name) => (
                                        <a href={providers[name]}>
                                            {name}
                                        </a>
                                    ))}
                                </div>
                            )}
                        </div>
                    )}
                </div>
            </div>
            <div className="footer">
                <a href="https://github.com/g-harel/targetblank">
                    <i className="fab fa-github" />
                </a>
            </div>
        </div>
    );
};

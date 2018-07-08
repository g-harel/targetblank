import "../static/page.home.scss";

import {api} from "../api";
import {input, props as inputP} from "../components/input";

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

const focusOnInput = () => {
    setTimeout(() => requestAnimationFrame(() => {
        const input: HTMLElement = document.querySelector("form.input input");
        if (input) {
            input.focus();
        }
    }));
};

export const home = (p, update) => {
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

    focusOnInput();

    return () => (
        ["div.home", {}, [
            ["div.header", {}, [
                ["i.far.fa-xs.fa-circle"],
                "targetblank",
            ]],
            ["div.screens", {className: {scrolled}}, [
                ["div.screen.signup", {}, [
                    [input, {
                        callback,
                        title: "Create a homepage",
                        placeholder: "john@example.com",
                        validator: /^\S+@\S+\.\S{2,}$/g,
                        message: "That doesn't look like an email address",
                    }],
                ]],
                ["div.screen.confirmation", {}, [
                    ["span.title", {}, [
                        ["i.far.fa-lg.fa-check-circle"],
                        "Confirmation Sent",
                    ]],
                    email && ["div.email", {}, [
                        ["a.address", {
                            href: email.link,
                            target: "_blank",
                            className: {inert: !email.link},
                        }, [
                            email.addr,
                            email.link && ["i.fas.fa-xs.fa-external-link-alt"],
                        ]],
                        !email.link && ["div.providers", {},
                            Object.keys(providers).map((p) => (
                                ["a", {href: providers[p].url}, [
                                    p,
                                ]]
                            )),
                        ],
                    ]],
                ]],
            ]],
            ["div.footer", {}, [
                ["a", {target: "_blank", href: "https://github.com/g-harel/targetblank"}, [
                    ["i.fab.fa-github"],
                ]],
            ]],
        ]]
    );
};

import "../static/page.home.scss";

import {Component, ComponentElement as CE} from "okwolo/lite";

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
    yahoo: {
        domains: ["yahoo"],
        url: "https://mail.yahoo.com",
    },
    outlook: {
        domains: ["live", "outlook", "hotmail", "msn"],
        url: "https://outlook.live.com",
    },
    icloud: {
        domains: ["icloud"],
        url: "https://mail.icould.com",
    },
};

const makeEmail = (addr: string): email => {
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

export const home: Component = (p, update) => {
    let email: email = null;

    const callback = async (addr: string) => {
        try {
            await api.page.create(addr);
        } catch (e) {
            console.log(e);
            return "Something went wrong";
        }

        email = makeEmail(addr);
        update();
    };

    const onclick = () => {
        email = email ? null : makeEmail("test@agmail.com");
        update();
    };

    focusOnInput();

    return () => (
        ["div.home", {}, [
            ["div.header", {onclick}, [
                ["i.far.fa-circle.fa-xs"],
                "targetblank",
            ]],
            ["div.screens", {className: {scrolled: !!email}}, [
                ["div.screen.signup", {}, [
                    <CE<inputP>>[input, {
                        callback,
                        title: "Create a homepage",
                        placeholder: "john@example.com",
                        validator: /^\S+@\S+\.\S{2,}$/g,
                        message: "That doesn't look like an email address",
                    }],
                ]],
                ["div.screen.confirmation", {}, [
                    ["span.title", {}, [
                        ["i.far.fa-check-circle"],
                        "Confirmation Sent",
                    ]],
                    email && ["div.email", {}, [
                        ["div.address", {}, [
                            email.addr,
                        ]],
                        ["div.link", {}, [
                            email.link ? (
                                ["a", {href: email.link}, [
                                    "visit your inbox",
                                ]]
                            ) : (
                                ["div.providers", {},
                                    Object.keys(providers).map((p) => {
                                        const {url} = providers[p];
                                        return (
                                            ["a", {href: url}, [
                                                p,
                                            ]]
                                        );
                                    }),
                                ]
                            ),
                        ]],
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

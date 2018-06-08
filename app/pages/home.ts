import "../static/page.home.scss";

import {Component, ComponentElement as CE} from "okwolo/lite";

import {api} from "../api";
import {input, props as inputP} from "../components/input";

export const home: Component = (p, update) => {
    setTimeout(() => requestAnimationFrame(() => {
        const input: HTMLElement = document.querySelector(".signup input");
        input.focus();
    }));

    let email: string = "";

    const callback = async (addr: string) => {
        try {
            await api.page.create(addr);
        } catch (e) {
            console.log(e);
            return "Something went wrong";
        }

        email = addr;
        update();
    };

    const onclick = () => {
        email = email ? "" : "test@gmail.com";
        update();
    };

    return () => (
        ["div.home", {}, [
            ["div.header", {onclick}, [
                ["i.far.fa-circle.fa-xs"],
                "targetblank",
            ]],
            ["div.mask", {}, [
                ["div.screens", {
                    className: {
                        scrolled: !!email,
                    },
                }, [
                    ["div.screen.signup", {}, [
                        <CE<inputP>>[input, {
                            callback,
                            title: "Create a homepage",
                            placeholder: "john@example.com",
                            validator: /^\S+@\S+\.\S+$/g,
                            message: "That doesn't look like an email address",
                        }],
                    ]],
                    ["div.screen", {}, [
                        email,
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

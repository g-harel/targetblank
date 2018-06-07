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

    const callback = (addr: string) => {
        console.log(addr);
        return new Promise((r) => setTimeout(() => r("Something went wrong"), 1500));
        email = addr;
        update();
    };

    return () => (
        ["div.home", {}, [
            ["div.header", {}, [
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
        ]]
    );
};

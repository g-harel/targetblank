import "../static/page.home.scss";

import {Component, ComponentElement as CE} from "okwolo/lite";

import {api} from "../api";
import {input, props as inputP} from "../components/input";

export const home: Component = () => {
    setTimeout(() => requestAnimationFrame(() => {
        const input: HTMLElement = document.querySelector(".signup input");
        input.focus();
    }));

    return () => (
        ["div.home", {}, [
            ["div.header", {}, [
                ["span.icon.fa-xs", {innerHTML:`
                    <i  class="fas fa-circle"
                        data-fa-mask="fas fa-circle"
                        data-fa-transform="shrink-7" />
                `}],
                "targetblank",
            ]],
            ["div.signup", {}, [
                <CE<inputP>>[input, {
                    placeholder: "john@example.com",
                    validator: /^\S+@\S+\.\S+$/g,
                    message: "Must be an email address.",
                    callback: console.log,
                }],
            ]],
        ]]
    );
};

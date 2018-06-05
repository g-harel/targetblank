import "../static/page.home.scss";

import {Component, ComponentElement} from "okwolo/lite";

import {api} from "../api";
import {input, props as inputProps} from "../components/input";

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
                <ComponentElement<inputProps>>[input, {
                    placeholder: "john@example.com",
                    validator: /^\S+@\S+\.\S+$/g,
                    message: "Must be an email address.",
                }],
            ]],
        ]]
    );
};

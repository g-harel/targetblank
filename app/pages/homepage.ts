import {Component} from "okwolo/lite";
import {api, IPageData} from "../api";

export interface IProps {
    addr: string;
    token?: string;
}

export const homepage: Component<IProps> = ({addr, token}, update) => {
    let data: IPageData | null = null;
    let error = null;

    api.page.fetch(addr, token)
        .then((d) => {
            data = d;
            update();
        })
        .catch((e) => {
            error = e;
            update();
        });

    return () => {
        if (error !== null) {
            return "login or an error has occurred";
        }
        if (data === null) {
            return "loading";
        }
        if (token) {
            // return "change your password";
        }
        return (
            ["pre | font-family: monospace; padding: 100px;", {}, [
                JSON.stringify(data, null, 2),
            ]]
        );
    };
};

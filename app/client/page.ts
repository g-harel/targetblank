import {app} from "../app";
import {component as error, show} from "./error";

const replaceToken = (str: string, name: string, value: string): string => {
    str = str.replace(`{{${name}?}}`, `(\?:\\/${value})\?`);
    return str.replace(`{{${name}}}`, `\\/${value}`);
};

export const page = <T>(
    path: string,
    ...paramList: string[]
) => (
    component: (params: T, update: any) => any,
) => {
    path = replaceToken(path, "addr", "(\\w{6})");
    path = replaceToken(path, "token", "([^\\s\\/]+)");
    path = replaceToken(path, "**", "(.*)");
    path = `^${path}\\/?$`;
    app(new RegExp(path), (params) => () => {
        paramList.forEach((item, index) => {
            params[item] = params[index];
        });
        return (
            ["div.page", {}, [
                [component, params],
                [error],
            ]]
        );
    });
};

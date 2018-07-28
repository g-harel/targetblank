import "../static/page.homepage.scss";

import {api, IError} from "../api";
import {read, save} from "../storage";

export interface IHomepageProps {
    addr: string;
}

const wrap = (element: any) => (
    ["div.homepage", {}, [
        element,
    ]]
);

export const homepage = ({addr}, update) => {
    let error: IError | null = null;
    let loaded = false;

    let {data, token} = read(addr);

    api.page.fetch(addr, token)
        .then((d) => {
            data = d;
            loaded = true;
            save(addr, {data});
            update();
        })
        .catch((e) => {
            error = e;
            console.log(e);
            update();
        });

    return () => (
        ["div.homepage", {}, [
            error ? "couldn't load" : null,
            console.log(data) || null,
            !loaded && (
                ["div.loading", {}, [
                    ["i.fa.fa-circle"],
                ]]
            ),
            !!data && (
                ["div.groups", {},
                    data.groups.map((group) => (
                        ["div.group", {}, [
                            group.meta.title || null,
                            ["div.items", {},
                                group.items.map((item) => (
                                    ["pre", {}, [JSON.stringify(item)]]
                                )),
                            ],
                        ]]
                    ))
                ]
            ),
        ]]
    );
};

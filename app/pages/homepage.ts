import "../static/page.homepage.scss";

import {api, IError} from "../client/api";
import {read, save} from "../client/storage";

export interface IHomepageProps {
    addr: string;
}

export const homepage = ({addr}, update) => {
    let error: IError | null = null;

    let {data, token} = read(addr);

    api.page.fetch(addr, token)
        .then((d) => {
            data = d;
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

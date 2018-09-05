import "../static/page.homepage.scss";

import {client, IPageData} from "../client/client";

export interface IHomepageProps {
    addr: string;
}

export const homepage = ({addr}, update) => {
    client.page.fetch(
        (data) => update(data, undefined),
        (err) => update(undefined, err),
        addr,
    );

    return (data?: IPageData, err?: string) => (
        ["div.homepage", {}, [
            err ? "couldn't load" : null,
            !data ? null : (
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
                    )),
                ]
            ),
        ]]
    );
};

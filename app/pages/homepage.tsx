import "../static/page.homepage.scss";

import {client, IPageData} from "../client/client";
import {PageComponent} from "../components/page";

export const Homepage: PageComponent = ({addr}, update) => {
    client.page.fetch(
        (data) => update(data, undefined),
        (err) => update(undefined, err),
        addr,
    );

    return (data?: IPageData, err?: string) => (
        <div className="homepage">
            {err && "couldn't load"}
            {data && (
                <div className="groups">
                    {data.groups.map((group) => (
                        <div className="group">
                            {group.meta.title || null}
                            <div className="items">
                                {group.items.map((item) => (
                                    <pre>
                                        {JSON.stringify(item)}
                                    </pre>
                                ))}
                            </div>
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
};

import {client, IPageData} from "../../library/client/client";
import {PageComponent} from "../../components/page";
import {styled} from "../../library/styled";

const Wrapper = styled("div")({});

const Groups = styled("div")({});

const Group = styled("div")({});

const Items = styled("div")({});

export const Homepage: PageComponent = ({addr}, update) => {
    client.page.fetch(
        (data) => update(data, undefined),
        (err) => update(undefined, err),
        addr,
    );

    return (data?: IPageData, err?: string) => (
        <Wrapper>
            {!!err && "couldn't load"}
            {!!data && (
                <Groups>
                    {data.groups.map((group) => (
                        <Group>
                            {group.meta.title || null}
                            <Items>
                                {group.items.map((item) => (
                                    <pre>
                                        {JSON.stringify(item)}
                                    </pre>
                                ))}
                            </Items>
                        </Group>
                    ))}
                </Groups>
            )}
        </Wrapper>
    );
};

import {client, IPageData} from "../../library/client/client";
import {PageComponent} from "../../components/page";
import {styled} from "../../library/styled";

const Wrapper = styled("div")({});

const Groups = styled("div")({});

const Group = styled("div")({});

const Items = styled("div")({});

interface Data {
    page?: IPageData;
    err?: string;
}

export const Homepage: PageComponent<Data> = ({addr}, update) => {
    client.page.fetch(
        (page) => update({page}),
        (err) => update({err}),
        addr,
    );

    return (data) => (
        <Wrapper>
            {!!data.err && "couldn't load"}
            {!!data.page && (
                <Groups>
                    {data.page.groups.map((group) => (
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

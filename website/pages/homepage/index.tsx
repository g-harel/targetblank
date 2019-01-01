import {client, IPageData} from "../../internal/client";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";

const Wrapper = styled("div")({});

const Groups = styled("div")({});

const Group = styled("div")({});

const Items = styled("div")({});

interface Data {
    page?: IPageData;
    err?: string;
}

export const Homepage: PageComponent<Data> = ({addr}, update) => {
    client.page.fetch((page) => update({page}), addr);

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
                                    <pre>{JSON.stringify(item)}</pre>
                                ))}
                            </Items>
                        </Group>
                    ))}
                </Groups>
            )}
        </Wrapper>
    );
};

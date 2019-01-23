import {app} from "../../internal/app";
import {client, IPageData} from "../../internal/client";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";

const Wrapper = styled("div")({});

const Groups = styled("div")({});

const Group = styled("div")({});

const Items = styled("div")({});

export const Document: PageComponent<IPageData> = ({addr}, update) => {
    const err = (message) => {
        console.warn(message);
        app.redirect(`/${addr}/login`);
    };

    client.page.fetch(update, err, addr);

    return (data: IPageData) => {
        // Response not yet received.
        if (!data) {
            // TODO
            return "loading";
        }

        return (
            <Wrapper>
                <Groups>
                    {...data.groups.map((group) => (
                        <Group>
                            {group.meta.title || ""}
                            <Items>
                                {...group.entries.map((item) => (
                                    <pre>{JSON.stringify(item)}</pre>
                                ))}
                            </Items>
                        </Group>
                    ))}
                </Groups>
            </Wrapper>
        );
    };
};

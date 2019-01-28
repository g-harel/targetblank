import {app} from "../../internal/app";
import {client, IPageData} from "../../internal/client";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";
import {Loading} from "../../components/loading";

const Wrapper = styled("div")({});

const EditButton = styled("div")({
    cursor: "pointer",
    display: "inline-block",
    float: "right",
    userSelect: "none",
    padding: "0.85rem 1.4rem",
    fontWeight: "bold",
});

const Groups = styled("div")({});

const Group = styled("div")({});

const Items = styled("div")({});

export const Document: PageComponent<IPageData> = ({addr}, update) => {
    client.page.fetch(update, () => app.redirect(`/${addr}/login`), addr);

    return (data: IPageData) => {
        // Response not yet received.
        if (!data) return <Loading />;

        return (
            <Wrapper>
                <EditButton onclick={() => app.redirect(`/${addr}/edit`)}>
                    edit
                </EditButton>
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

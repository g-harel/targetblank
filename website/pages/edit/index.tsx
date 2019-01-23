import {app} from "../../internal/app";
import {client, IPageData} from "../../internal/client";
import {PageComponent} from "../../components/page";
import {styled} from "../../internal/styled";

const editorID = "targetblank-editor";
const headerHeight = "2.9rem";
const saveDelay = 1000;

const Wrapper = styled("div")({});

const Header = styled("header")({
    borderBottom: "1px solid #ddd",
    height: headerHeight,
    padding: "0.85rem 1.4rem",
    userSelect: "none",
});

const BackButton = styled("div")({
    cursor: "pointer",
    display: "inline-block",
    userSelect: "none",
    fontWeight: "bold",
});

const Status = styled("div")({
    display: "inline-block",
    float: "right",
    fontWeight: "bold",

    "&.error": {
        color: "tomato",
    },
});

const SavedIcon = styled("i")({
    color: "yellowgreen",
    margin: "0 0.4rem",
    transform: "translate(0, 0.1rem)",
});

const Editor = styled("textarea")({
    border: "none",
    fontFamily: "Inconsolata, monospace",
    fontSize: "1.2rem",
    height: `calc(100% - ${headerHeight})`,
    lineHeight: "1.3",
    padding: "2.2rem 0 1rem 2.2rem",
    outline: "none",
    resize: "none",
    whiteSpace: "pre",
    width: "100%",
});

interface Data {
    page: IPageData;
    status: "saving" | "saved" | "error";
    error?: string;
}

export const Edit: PageComponent<Data> = ({addr}, update) => {
    const callback = (data: IPageData) => {
        update({
            page: data,
            status: "saved",
        });
    };

    const err = (message) => {
        console.warn(message);
        app.redirect(`/${addr}/login`);
    };

    client.page.fetch(callback, err, addr);

    return (data?: Data) => {
        // Response not yet received.
        if (!data) {
            // TODO
            return "loading";
        }

        let statusContent: any = null;
        if (data.status === "error") {
            statusContent = data.error;
        } else if (data.status === "saving") {
            statusContent = "saving ...";
        } else if (data.status === "saved") {
            statusContent = (
                <span>
                    <SavedIcon className="far fa-lg fa-check-circle" />
                    saved
                </span>
            );
        }

        let timeout: any = null;
        const onChange = (e) => {
            const {value} = e.target;
            clearTimeout(timeout);
            timeout = setTimeout(() => {
                update({page: data.page, status: "saving"});

                const callback = () => {
                    update({page: data.page, status: "saved"});
                };
                const err = (message) => {
                    update({page: data.page, status: "error", error: message});
                };
                client.page.edit(callback, err, addr, value);
            }, saveDelay);
        };

        return (
            <Wrapper>
                <Header>
                    <BackButton onclick={() => app.redirect(`/${addr}`)}>
                        back
                    </BackButton>
                    <Status className={{error: !!data.error}}>
                        {statusContent}
                    </Status>
                </Header>
                <Editor
                    id={editorID}
                    value={data.page.raw}
                    oninput={onChange}
                />
            </Wrapper>
        );
    };
};

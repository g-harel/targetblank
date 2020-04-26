import {Input} from "../../components/input";
import {styled} from "../../internal/style";
import {color, size, font} from "../../internal/style/theme";
import {Header} from "../../components/header";
import {PageComponent} from "../../components/page";
import {write, ExtensionStore, read} from "../../internal/extension";
import {routes, path, safeRedirect} from "../../routes";
import {Anchor} from "../../components/anchor";

const Wrapper = styled("div")({
    alignItems: "center",
    display: "flex",
    flexDirection: "column",
});

const Info = styled("div")({
    color: color.textSecondarySmall,
    fontSize: size.tiny,
    fontWeight: 600,
    marginTop: "-0.5rem",
    textAlign: "center",
    width: "18rem",
});

const Current = styled("div")({
    alignItems: "center",
    display: "flex",
    flexDirection: "column",
    fontSize: size.tiny,
    marginTop: "3em",
});

const Address = styled("div")({
    border: "1px solid black",
    borderColor: color.decoration,
    borderRadius: "3px",
    display: "block",
    fontFamily: font.monospace,
    fontSize: size.normal,
    margin: "0.25em",
    padding: "0.5em 1em",
});

export const Options: PageComponent = (_, update) => {
    document.title = "targetblank - options";
    let options: ExtensionStore | null = null;

    const submit = async (addr: string) => {
        await write({addr});
        setTimeout(() => {
            safeRedirect(routes.document, addr);
        });
        return "";
    };

    read().then((opts) => {
        options = opts;
        update();
    });

    return () => (
        <Wrapper>
            <Header muted sub="extension options" />
            <Input
                title="set homepage address"
                callback={submit}
                validator={/^\w{6}$|^local$/g}
                message="invalid address"
                placeholder="a1b2c3"
                focus
            />
            <Info>
                Use the "local" address to only store your page on this
                computer. Document must still be parsed on the server.
            </Info>
            {options && !!options.addr && (
                <Current>
                    current address
                    <Anchor
                        id="document"
                        href={path(routes.document, options.addr)}
                    >
                        <Address>{options.addr}</Address>
                    </Anchor>
                </Current>
            )}
        </Wrapper>
    );
};

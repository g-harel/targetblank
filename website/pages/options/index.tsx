import {Input} from "../../components/input";
import {styled, colors, size, fonts} from "../../internal/style";
import {Header} from "../../components/header";
import {PageComponent} from "../../components/page";
import {write, ExtensionStore, read} from "../../internal/extension";
import {routes, path, safeRedirect} from "../../routes";
import {Anchor} from "../../components/anchor";

const Wrapper = styled("div")({});

const Current = styled("div")({
    alignItems: "center",
    display: "flex",
    flexDirection: "column",
    fontSize: size.tiny,
    marginTop: "3em",
});

const Address = styled("div")({
    border: "1px solid black",
    borderColor: colors.decoration,
    borderRadius: "3px",
    display: "block",
    fontFamily: fonts.monospace,
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
            <Header muted />
            <Input
                title="set homepage address"
                callback={submit}
                validator={/^\w{6}$|^local$/g}
                message="invalid address"
                placeholder="a1b2c3"
                focus
            />
            {options && !!options.addr && (
                <Current>
                    current address
                    <Anchor href={path(routes.document, options.addr)}>
                        <Address>{options.addr}</Address>
                    </Anchor>
                </Current>
            )}
        </Wrapper>
    );
};

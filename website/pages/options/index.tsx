import {Input} from "../../components/input";
import {styled, colors, size, fonts} from "../../internal/style";
import {Header} from "../../components/header";
import {PageComponent} from "../../components/page";

const Wrapper = styled("div")({});

const Current = styled("div")({
    alignItems: "center",
    display: "flex",
    flexDirection: "column",
    fontSize: size.tiny,
    paddingTop: "3em",
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

export const Options: PageComponent = ({addr}) => () => {
    document.title = "targetblank - options";

    const submit = (pass: string) => {
        return new Promise<string>((resolve) => {
            console.log(pass);
            resolve("");
        });
    };

    return (
        <Wrapper>
            <Header muted />
            <Input
                title="set homepage address"
                callback={submit}
                validator={/.*/g}
                message="invalid address"
                placeholder="a1b2c3"
                focus
            />
            {!!addr && (
                <Current>
                    current address
                    <Address>
                        {addr}
                    </Address>
                </Current>
            )}
        </Wrapper>
    );
};

// The h function provided by okwolo is attached to the global object.
import h from "okwolo/src/h";
(window as any).h = h;

import "../website/global.css";
import "normalize.css";

import {app} from "../website/internal/app";
import {Input} from "../website/components/input";
import {styled, colors, size, fonts} from "../website/internal/style";
import {Header} from "../website/components/header";

app.use("target", document.body);

app.setState({});

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

app(() => () => {
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
            <Current>
                current address
                <Address>
                    a1b2c3
                </Address>
            </Current>
        </Wrapper>
    );
});

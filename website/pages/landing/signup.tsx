import {styled} from "../../internal/style";
import {Component} from "../../internal/types";
import {Input, Props as InputProps} from "../../components/input";

const Wrapper = styled("div")({});

export interface Props {
    callback: InputProps["callback"];
}

export const Signup: Component<Props> = (props) => () => (
    <Wrapper>
        <Input
            callback={props.callback}
            title="create a page"
            type="email"
            autocomplete="email"
            placeholder="email@example.com"
            validator={/^\S+@\S+\.\S{2,}$/g}
            message="invalid email address"
            focus={true}
        />
    </Wrapper>
);

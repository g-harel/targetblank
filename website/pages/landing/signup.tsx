import {styled} from "../../internal/styled";
import {Component} from "../../internal/types";
import {Input, Props as InputProps} from "../../components/input";

const Wrapper = styled("div")({
    opacity: 0,

    "&.visible": {
        opacity: 1,
    },
});

export interface Props {
    callback: InputProps["callback"];
    visible?: boolean;
}

export const Signup: Component<Props> = (props) => () => (
    <Wrapper className={{visible: props.visible}}>
        <Input
            callback={props.callback}
            title="create a page"
            type="email"
            placeholder="john@example.com"
            validator={/^\S+@\S+\.\S{2,}$/g}
            message="That doesn't look like an email address"
            focus={true}
        />
    </Wrapper>
);

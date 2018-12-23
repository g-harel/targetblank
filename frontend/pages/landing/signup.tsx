import {styled} from "../../library/styled";

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

export const Signup = (props: Props) => () => (
    <Wrapper className={{visible: props.visible}}>
        <Input
            callback={props.callback}
            title="Create a homepage"
            placeholder="john@example.com"
            validator={/^\S+@\S+\.\S{2,}$/g}
            message="That doesn't look like an email address"
            focus={true}
        />
    </Wrapper>
);

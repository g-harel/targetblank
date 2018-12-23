import {Input, Props as InputProps} from ".";
import {Component} from "../../library/types";

export interface Props {
    title: string;
    callback: InputProps["callback"];
}

export const Password: Component<Props> = ({title, callback}) => () => (
    <Input
        title={title}
        callback={callback}
        type="password"
        validator={/.{8,32/ig}
        message="Password is too short"
        placeholder="password123"
        focus={true}
    />
);

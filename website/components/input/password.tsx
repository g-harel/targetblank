import {Input, Props as InputProps} from ".";
import {Component} from "../../internal/types";

export interface Props {
    title: string;
    hint?: string;
    autocomplete?: string;
    validate?: boolean;
    callback: InputProps["callback"];
}

export const Password: Component<Props> = ({
    title,
    hint,
    autocomplete,
    callback,
    validate,
}) => () => (
    <Input
        title={title}
        hint={hint}
        callback={callback}
        type="password"
        autocomplete={autocomplete}
        validator={validate ? /.{8,32}/gi : /./}
        message="password is too short"
        placeholder="password123"
        focus={true}
    />
);

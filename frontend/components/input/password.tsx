import {Input, Props as InputProps} from ".";

export interface Props {
    title: string;
    callback: InputProps["callback"];
}

export const Password = ({title, callback}: Props) => () => (
    <div className="password-input">
        <Input
            title={title}
            callback={callback}
            type="password"
            validator={/.{8,32/ig}
            message="Password is too short"
            placeholder="password123"
            focus={true}
        />
    </div>
);

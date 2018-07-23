import {input, IInputProps} from "./input";

export interface IPasswordProps {
    title: string;
    callback: IInputProps["callback"];
}

export const password = ({title, callback}: IPasswordProps) => () => (
    ["div.password-input", {} , [
        [input, {
            title,
            callback,
            type: "password",
            validator: /.{8,32}/ig,
            message: "Password is too short",
            placeholder: "password123",
            focus: true,
        }],
    ]]
);

export interface IErrorProps {
    errors: string[];
    hide: () => void;
}

export const errors = ({errors}: IErrorProps) => () => (
    ["div.errors", {} ,
        errors.map((err) => (
            ["div.error", {}, [
                err,
            ]]
        ))
    ]
);

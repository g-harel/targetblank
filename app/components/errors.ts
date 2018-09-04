import "../static/component.errors.scss";

export interface IErrorProps {
    errors: string[];
    hide: () => void;
}

export const errors = ({errors, hide}: IErrorProps) => () => (
    ["div.errors", {
        className: {hidden: !errors.length},
    }, [
        ["div.title", {}, [
            "error",
            ["div.dismiss", {onclick: hide}, [
                "dismiss",
            ]],
        ]],
        ...errors.map((err) => (
            ["div.error", {}, [
                err,
            ]]
        )),
    ]]
);

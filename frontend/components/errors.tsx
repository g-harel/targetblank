import "../static/component.errors.scss";

export interface Props {
    errors: string[];
    hide: () => void;
}

export const Errors = ({errors, hide}: Props) => () => (
    <div className={{
        errors: true,
        hidden: !errors.length,
    }}>
        <div className="title">
            error
            <div className="dismiss" onclick={hide}>
                dismiss
            </div>
        </div>
        {...errors.map((err) => (
            <div className="error">
                {err}
            </div>
        ))}
    </div>
);

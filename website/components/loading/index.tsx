import {Component} from "../../internal/types";
import {Header} from "../header";
import {Icon} from "../icon";
import {styled} from "../../internal/style";

const Spinner = styled("div")({
    opacity: 0.2,
    textAlign: "center",
    padding: "2rem",
});

export const Loading: Component = () => () => {
    return (
        <div>
            <Header muted />
            <Spinner>
                <Icon name="spinner" spin size={1.5} />
            </Spinner>
        </div>
    );
};

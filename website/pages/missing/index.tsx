import {Component} from "../../internal/types";
import {Header} from "../../components/header";

export const Missing: Component<{}> = () => () => <Header muted title="404" />;

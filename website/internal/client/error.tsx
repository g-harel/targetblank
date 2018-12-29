import {Errors} from "../../components/errors";
import {Component} from "../../internal/types";

let errList: string[] = [];
let updator: null | Function = null;

export const show = (...errs: string[]) => {
    errs.forEach((err) => console.error(err));
    errList.push(...errs);
    update();
};

const hide = () => {
    errList = [];
    update();
};

const update = () => {
    if (updator) {
        updator();
    }
};

export const ErrorComponent: Component = (_, u) => {
    updator = u;
    return () => (
        <Errors hide={hide} errors={errList} />
    );
};

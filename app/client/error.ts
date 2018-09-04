import {errors, IErrorProps} from "../components/errors";

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

export const component = (_, u) => {
    updator = u;
    return () => (
        [errors, {
            hide,
            errors: errList,
        } as IErrorProps]
    );
};

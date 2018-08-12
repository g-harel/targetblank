import {errors, IErrorProps} from "../components/errors";

let errList: string[] = [];
let updator: null | Function = null;

export const show = (...errs: string[]) => {
    errList.push(...errs);
    update();
}

const hide = () => {
    errList = [];
    update();
}

const update = () => {
    if (updator) {
        updator();
    }
}

export const component = (_, u) => {
    updator = u;
    return () => (
        [errors, {
            errors: errList,
            hide,
        } as IErrorProps]
    )
}

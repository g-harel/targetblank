import {showChip} from "../components/page/chips";
import {safeRedirect, routes} from "../routes";
import {IRequestError} from "./api";

export const genRequestErrHandler = (addr: string) => (err: IRequestError) => {
    if (err.isAuth) {
        handleAuthErr(addr);
    } else {
        handleErr();
    }
};

export const handleAuthErr = (addr: string) => {
    showChip("not authorized", 4000);
    safeRedirect(routes.login, addr!);
};

export const handleErr = (resolve?: () => void) => {
    showChip("something went wrong", 4000);
    if (resolve) resolve();
};

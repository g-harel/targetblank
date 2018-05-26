declare module "okwolo/lite" {
    export type State = Exclude<any, undefined>;

    export interface App {
        (any): void

        setState(state: State): void
        setState(updater: (state: State) => State): void

        getState(): State
    }

    export default function(target?: any, global?: any): App
}

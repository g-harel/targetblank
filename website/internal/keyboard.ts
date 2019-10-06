export interface Event {
    ctrl: boolean;
    key: KeyboardEvent["key"];
    shift: boolean;
}

export type Handler = (e: Event) => any;

let handlers: Handler[] = [];

window.onkeydown = (e: KeyboardEvent) => {
    const event: Event = {
        ctrl: navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey,
        key: e.key,
        shift: e.shiftKey,
    };
    handlers.forEach((h) => h(event));
};

// Adds a new handlers.
export const keyboard = (h: Handler) => {
    handlers.push(h);
};

// Remove all existing handlers.
export const reset = () => {
    handlers = [];
};

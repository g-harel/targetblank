export interface Event {
    key: KeyboardEvent["key"];
    ctrl: boolean;
}

export type Handler = (e: Event) => any;

let handlers: Handler[] = [];

window.onkeydown = (e: KeyboardEvent) => {
    const event: Event = {
        key: e.key,
        ctrl: navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey,
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

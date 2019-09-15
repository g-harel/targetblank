const storageKey = "targetblank-storage";

export const isExtension = !!(
    (window as any).chrome &&
    (window as any).chrome.runtime &&
    (window as any).chrome.runtime.id
);
export const isChromeExtension = !window.browser && isExtension;

export interface ExtensionStore {
    addr: string | null;
}

// Generates a zeroed-out page storage value.
const empty = (): ExtensionStore => ({addr: null});

// Wrapper around browser's `storage.sync.get` to support both Chrome and Firefox.
const crossRead = (keys: any): Promise<any> => {
    if (isChromeExtension) {
        return new Promise((resolve) => {
            (window as any).chrome.storage.sync.get(keys, resolve);
        });
    }
    return browser.storage.sync.get(keys);
};

// Wrapper around browser's `storage.sync.set` to support both Chrome and Firefox.
const crossWrite = (keys: browser.storage.StorageObject): Promise<void> => {
    if (isChromeExtension) {
        return new Promise((resolve) => {
            (window as any).chrome.storage.sync.set(keys);
            resolve();
        });
    }
    return browser.storage.sync.set(keys);
};

// Reads extension data from synced browser storage.
export const read = async (): Promise<ExtensionStore> => {
    if (!isExtension) return empty();
    const res = await crossRead(storageKey);
    return (res as any)[storageKey] || empty();
};

// Updates extension data in synced browser storage.
export const write = async (values: Partial<ExtensionStore>) => {
    const current = await read();
    const updated = Object.assign(current, values);
    await crossWrite({[storageKey]: updated});
};

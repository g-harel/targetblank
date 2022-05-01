const storageKey = "targetblank-storage";
const fallbackLocalStorageKey = "targetblank-storage-key";

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

// Check if storage permission had been granted.
const hasStoragePermission = async (): Promise<boolean> => {
    if (!isExtension) return false;
    return new Promise((res) => {
        (window as any).chrome.permissions.contains(
            {
                permissions: ["storage"],
            },
            (result: any) => res(!!result),
        );
    });
};

// Wrapper around browser's `storage.sync.get` to support both Chrome and Firefox.
const crossRead = async (keys: any): Promise<any> => {
    if (!(await hasStoragePermission())) {
        return JSON.parse(
            localStorage.getItem(fallbackLocalStorageKey) || "{}",
        );
    }
    if (isChromeExtension) {
        return new Promise((resolve) => {
            (window as any).chrome.storage.sync.get(keys, resolve);
        });
    }
    return browser.storage.sync.get(keys);
};

// Wrapper around browser's `storage.sync.set` to support both Chrome and Firefox.
const crossWrite = async (
    keys: browser.storage.StorageObject,
): Promise<void> => {
    if (!(await hasStoragePermission())) {
        localStorage.setItem(fallbackLocalStorageKey, JSON.stringify(keys));
    }
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
    await crossWrite({[storageKey]: updated} as Record<string, any>);
};

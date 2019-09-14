const storageKey = "targetblank-storage";

export const isExtension =
    (window as any).browser &&
    (window as any).browser.runtime &&
    (window as any).browser.runtime.id;

export interface ExtensionStore {
    addr: string | null;
}

// Generates a zeroed-out page storage value.
const empty = (): ExtensionStore => ({addr: null});

// Reads extension data from synced browser storage.
export const read = async (): Promise<ExtensionStore> => {
    if (!isExtension) return empty();
    const res = await browser.storage.sync.get(storageKey);
    return (res as any)[storageKey] || empty();
};

// Updates extension data in synced browser storage.
export const write = async (values: Partial<ExtensionStore>) => {
    const current = await read();
    const updated = Object.assign(current, values);
    await browser.storage.sync.set({[storageKey]: updated});
};

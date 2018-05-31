import request from "request-promise-native";

const apiEndpoint = "https://api.targetblank.org";

export const fetchPage = async (token: string, address: string): Promise<IPageData> => {
    return request({
        uri: `${apiEndpoint}/page/${address}`,
        headers: {
            Token: token,
        },
        json: true,
    });
};

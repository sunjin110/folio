// import { google } from "googleapis";

export const OPTIONS = {
    clientId: process.env.GOOGLE_OAUTH_CLIENT_ID,
    clientSecret: process.env.GOOGLE_OAUTH_SECRET_ID,
    redirectUri: process.env.GOOGLE_OAUTH_REDIRECT_URL,
};

// export function createOAuth2Client(options?: {clientId?: string, clientSecret?: string, redirectUri?: string}) {
//     const { clientId, clientSecret, redirectUri } = {
//         ...OPTIONS,
//         ...options,
//     };

//     return new google.auth.OAuth2(clientId, clientSecret, redirectUri);
// }

interface QueryParams {
    client_id?: string;
    redirect_uri?: string;
    response_type?: string;
    scope?: string; // スペース区切り
    access_type?: string;
}

// export function generateOAuth2Url(): string {
//     return "test";
// }

export function generateOAuth2Url(): string {
    const baseUrl = process.env.REACT_APP_GOOGLE_OAUTH_URL; // OAuth URLを環境変数から取得
    console.log("process.env is ", process.env);
    const queryParams: QueryParams = {
        client_id: process.env.REACT_APP_GOOGLE_OAUTH_CLIENT_ID,
        redirect_uri: process.env.REACT_APP_GOOGLE_OAUTH_REDIRECT_URL,
        response_type: 'code',
        scope: 'profile email',
        access_type: 'offline', // TODO localのみ
    };

    // クエリパラメータをエンコードして組み立てる
    const queryString = Object.keys(queryParams)
        .filter(key => queryParams[key as keyof QueryParams] !== undefined)
        .map(key => {
            const value = queryParams[key as keyof QueryParams];
            return `${encodeURIComponent(key)}=${encodeURIComponent(value!)}`;
        })
        .join('&');

    // 完全なURLを生成
    return `${baseUrl}?${queryString}`;
}

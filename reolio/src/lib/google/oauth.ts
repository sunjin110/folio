import { google } from "googleapis";

export const OPTIONS = {
    clientId: process.env.GOOGLE_OAUTH_CLIENT_ID,
    clientSecret: process.env.GOOGLE_OAUTH_SECRET_ID,
    redirectUri: process.env.GOOGLE_OAUTH_REDIRECT_URL,
};

export function createOAuth2Client(options?: {clientId?: string, clientSecret?: string, redirectUri?: string}) {
    const { clientId, clientSecret, redirectUri } = {
        ...OPTIONS,
        ...options,
    };

    return new google.auth.OAuth2(clientId, clientSecret, redirectUri);
}

interface QueryParams {
    client_id?: string;
    redirect_uri?: string;
    response_type?: string;
    scope?: string; // スペース区切り
    access_type?: string;
}

// export function generateOAuth2Url(): string {
//     const queryParams: QueryParams = {
//         client_id: process.env.GOOGLE_OAUTH_CLIENT_ID,
//         redirect_uri: process.env.GOOGLE_OAUTH_REDIRECT_URL,
//         response_type: 'code',
//         scope: 'profile email',
//         access_type: 'offline', // TODO localのみ
//     };

//     const url = new URL(process.env.GOOGLE_OAUTH_URL);

//     Object.keys(queryParams).forEach(key => {
//         const value = queryParams[key as keyof QueryParams];
//         if (value !== undefined) {
//             url.searchParams.append(key, value.toString());
//         }
//     });

//     return url.toString();
// }

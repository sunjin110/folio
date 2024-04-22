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

// const generateOAuth2Url:  = () => {
//     return "";
// };

// https://accounts.google.com/o/oauth2/v2/auth?client_id=682633467318-vvlia00uaag3jplkls0uj1md371k54as.apps.googleusercontent.com&
// redirect_uri=http://localhost:3001/auth/google-oauth/callback&
// response_type=code&scope=profile email&access_type=offline

interface QueryParams {
    client_id?: string;
    redirect_uri?: string;
    response_type?: string;
    scope?: string; // スペース区切り
    access_type?: string;
}

export function generateOAuth2Url(): string {
    const queryParams: QueryParams = {
        client_id: process.env.GOOGLE_OAUTH_CLIENT_ID,
        redirect_uri: process.env.GOOGLE_OAUTH_REDIRECT_URL,
        response_type: 'code',
        scope: 'profile email',
        access_type: 'offline', // TODO localのみ
    };

    const url = new URL(process.env.GOOGLE_OAUTH_URL);

    Object.keys(queryParams).forEach(key => {
        const value = queryParams[key as keyof QueryParams];
        if (value !== undefined) {
            url.searchParams.append(key, value.toString());
        }
    });

    return url.toString();
}

//   // HTTP GET リクエストを行う非同期関数
//   async function fetchResource(url: string, queryParams: QueryParams): Promise<any> {
//     const urlObject = new URL(url);
  
//     // クエリパラメータを動的に追加
//     Object.keys(queryParams).forEach(key => {
//       const value = queryParams[key as keyof QueryParams];
//       if (value !== undefined) { // undefined のパラメータは追加しない
//         urlObject.searchParams.append(key, value.toString());
//       }
//     });
  
//     try {
//       const response = await fetch(urlObject.toString());
//       if (!response.ok) {
//         throw new Error('Network response was not ok');
//       }
//       return response.json(); // レスポンスボディをJSONとしてパース
//     } catch (error) {
//       console.error('Failed to fetch resource:', error);
//       throw error;
//     }
//   }
  
//   // 関数の使用例
//   fetchResource('https://jsonplaceholder.typicode.com/posts', {
//     userId: 10,
//     limit: 5,
//     sortBy: 'date',
//     asc: true
//   }).then(data => console.log(data));
  
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

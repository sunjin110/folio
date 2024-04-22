
export const OPTIONS = {
    clientId: process.env.GOOGLE_OAUTH_CLIENT_ID,
    clientSecret: process.env.GOOGLE_OAUTH_SECRET_ID,
};

export function createOAuth2Client(options?: {clientId?: string, clientSecret?: string, redirectUri?: string}) {

}
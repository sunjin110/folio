
export const OPTIONS = {
    clientId: process.env.GOOGLE_OAUTH_CLIENT_ID,
    clientSecret: process.env.GOOGLE_OAUTH_SECRET_ID,
    redirectUrl: process.env.GOOGLE_OAUTH_REDIRECT_URL,
};

export function createOAuth2Client(options?: {clientId?: string, clientSecret?: string, redirectUri?: string}) {

}

declare module 'process' {
    global {
        namespace NodeJS {
            interface ProcessEnv {
                GOOGLE_OAUTH_CLIENT_ID: string;
                GOOGLE_OAUTH_SECRET_ID: string;
                GOOGLE_OAUTH_REDIRECT_URL: string;
            }
        }
    }
}

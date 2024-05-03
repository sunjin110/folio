declare module "process" {
  global {
    namespace NodeJS {
      interface ProcessEnv {
        REACT_APP_GOOGLE_OAUTH_CLIENT_ID: string;
        REACT_APP_GOOGLE_OAUTH_SECRET_ID: string;
        REACT_APP_GOOGLE_OAUTH_REDIRECT_URL: string;
        REACT_APP_GOOGLE_OAUTH_URL: string;
        REACT_APP_GOLIO_BASE_URL: string;
      }
    }
  }
}

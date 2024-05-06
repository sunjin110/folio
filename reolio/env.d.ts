declare module "process" {
  global {
    namespace NodeJS {
      interface ProcessEnv {
        REACT_APP_GOLIO_BASE_URL: string;
      }
    }
  }
}

import { AuthError, InternalError } from "./error";

interface Toast {
  title?: string;
  description?: string;
}

interface HandlerErrorResponse {
  toast: Toast;
  navigationPath?: string;
}

export function handleError(err: unknown): HandlerErrorResponse {
  console.error(err);
  if (err instanceof AuthError) {
    return {
      toast: {
        title: "🔑 please login again 🔑",
        description: err.message,
      },
      navigationPath: "/login",
    };
  } else if (err instanceof InternalError) {
    return {
      toast: {
        title: "🐶 server error 🐶",
        description: err.message,
      },
    };
  }
  return {
    toast: {
      title: "⛔️ internal error ⛔️",
      description: `${err}`,
    },
  };
}

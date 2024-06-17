import { AuthError, InternalError } from "@/error/error";
import { ResponseError } from "@/generate/schema/http";

export function handleError(err: unknown, msg: string) {
  console.error("infrastructure/repository", err, msg);
  if (err instanceof ResponseError) {
    if (err.response.status === 401 || err.response.status === 403) {
      throw new AuthError(msg, err);
    }
    throw new InternalError(msg, err);
  }
  throw new InternalError(msg, err);
}

export function wrapError(err: unknown, msg: string): Error {
  console.error("infrastructure/repository", err, msg);
  if (err instanceof ResponseError) {
    if (err.response.status === 401 || err.response.status === 403) {
      return new AuthError(msg, err);
    }
    return new InternalError(msg, err);
  }
  return new InternalError(msg, err);
}

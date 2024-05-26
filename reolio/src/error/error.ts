export class FolioBaseError extends Error {
    name: string;
    stack?: string | undefined;
    cause?: unknown;

    constructor(message: string, cause?: unknown) {
        super(message);

        this.name = "FolioBaseError";
        this.cause = cause;

        if (Error.captureStackTrace) {
            Error.captureStackTrace(this, FolioBaseError);
        }
    }
}

export class InternalError extends FolioBaseError {
    constructor(message: string, cause?: unknown) {
        super(message);
        this.name = "InternalError";
        this.cause = cause;
    }
}

export class AuthError extends FolioBaseError {
    constructor(message: string, cause?: unknown) {
        super(message);
        this.name = "AuthError";
        this.cause = cause;
    }
}

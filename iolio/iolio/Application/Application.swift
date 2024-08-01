public enum Application {}

extension Application {
    struct ApplicationError: Error {
        enum ErrorKind {
            case internalError
        }
        
        let message: String
        let innerError: Error?
        let kind: ErrorKind
    }
}

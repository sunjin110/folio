public enum Usecase {}

extension Usecase {
    struct UsecaseError: Error {
        enum ErrorKind {
            case internalError
        }
        
        let message: String
        let innerError: Error?
        let kind: ErrorKind
    }
}

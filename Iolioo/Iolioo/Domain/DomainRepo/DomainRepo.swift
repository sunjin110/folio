public enum DomainRepo {}

public extension DomainRepo {
    struct RepoError: Error {
        enum ErrorKind {
            case internalError
        }
        
        let message: String
        let innerError: Error?
        let kind: ErrorKind
    }
}

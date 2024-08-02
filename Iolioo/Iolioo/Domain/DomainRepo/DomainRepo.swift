public enum DomainRepo {}

extension DomainRepo {
    public struct RepoError: Error {
        enum ErrorKind {
            case internalError
        }

        let message: String
        let innerError: Error?
        let kind: ErrorKind
    }
}

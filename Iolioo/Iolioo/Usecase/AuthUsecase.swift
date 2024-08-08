extension Usecase {
    protocol AuthUsecase {
        func verifyTokenAndStartSession(idToken: String, accessToken: String, refreshToken: String) async -> Result<Bool, UsecaseError>
    }
}

extension Usecase {
    class AuthUsecaseImpl: AuthUsecase {
        var authRepo: DomainRepo.Auth
        init(authRepo: DomainRepo.Auth) {
            self.authRepo = authRepo
        }
        
        func verifyTokenAndStartSession(idToken: String, accessToken: String, refreshToken: String) async -> Result<Bool, Usecase.UsecaseError> {
            let result = await self.authRepo.verifyTokenAndStartSession(idToken: idToken, accessToken: accessToken, refreshToken: refreshToken)
            
            switch result {
            case .success(let ok):
                return .success(ok)
            case .failure(let err):
                return .failure(.init(message: "failed authRepo.verifyTokenAndStartSession", innerError: err, kind: .internalError))
            }
        }
    }
}

#if DEBUG
extension Usecase {
    struct AuthUsecaseMock: AuthUsecase {
        var verifyTokenAndStartSessionResult: Result<Bool, UsecaseError>?
        func verifyTokenAndStartSession(idToken: String, accessToken: String, refreshToken: String) async -> Result<Bool, Usecase.UsecaseError> {
            guard let result = self.verifyTokenAndStartSessionResult else {
                return .failure(.init(message: "unspecified", innerError: nil, kind: .internalError))
            }
            return result
        }
    }
}
#endif

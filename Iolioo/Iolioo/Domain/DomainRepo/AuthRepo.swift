extension DomainRepo {
    protocol Auth {
        func verifyTokenAndStartSession(idToken: String, accessToken: String, refreshToken: String)
            async -> Result<Bool, DomainRepo.RepoError>
    }
}

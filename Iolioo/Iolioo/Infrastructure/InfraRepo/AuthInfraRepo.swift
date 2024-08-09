import Foundation

extension InfraRepo {
    class Auth: DomainRepo.Auth {
        
        let urlSession = URLSession(configuration: .default)
        let baseUrl: String
        let verifyTokenAndStartSessionApiPath = "/auth/google-oauth/verify-token-and-start-session"
        
        init(baseUrl: String) {
            self.baseUrl = baseUrl
        }
        
        func verifyTokenAndStartSession(idToken: String, accessToken: String, refreshToken: String) async -> Result<Bool, DomainRepo.RepoError> {
            guard var components = URLComponents(string: "\(self.baseUrl)\(self.verifyTokenAndStartSessionApiPath)") else {
                return .failure(.init(message: "failed make compoents. baseUrl: \(baseUrl), path: \(verifyTokenAndStartSessionApiPath)", innerError: nil, kind: .internalError))
            }
            
            components.queryItems = [
              URLQueryItem(name: "id_token", value: idToken),
              URLQueryItem(name: "access_token", value: accessToken),
              URLQueryItem(name: "refresh_token", value: refreshToken)
            ]
            
            guard let url = components.url else {
                return .failure(.init(message: "invalid url", innerError: nil, kind: .internalError))
            }
            
            var req = URLRequest(url: url)
            req.httpMethod = "GET"
            
            var resp: URLResponse?
            do {
                (_, resp) = try await self.urlSession.data(for: req)
            } catch {
                return .failure(.init(message: "failed request", innerError: error, kind: .internalError))
            }
            
            guard let resp = resp as? HTTPURLResponse else {
                return .failure(.init(message: "failed ", innerError: nil, kind: .internalError))
            }
            
            if resp.statusCode != 200 {
                return .failure(.init(message: "failed verify token and start session. statusCode: \(resp.statusCode)", innerError: nil, kind: .internalError))
            }
            return .success(true)
        }
    }
}

import OpenAPIURLSession
import SwiftData
import SwiftUI
import GoogleSignIn

@main
struct IoliooApp: App {
    let authUsecase: Usecase.AuthUsecase
    let articleUsecase: Usecase.ArticleUsecase
    
    internal init() {
        let appConf = AppConfiguration.shared

        print("appConf is \(appConf)")
        
        let serverURL = Foundation.URL(string: appConf.golioApiUrl)!
        let golioClient: APIProtocol = Client(
            serverURL: serverURL, transport: URLSessionTransport())
        let articleRepo = InfraRepo.Article(client: golioClient)
        let authRepo = InfraRepo.Auth(baseUrl: appConf.golioApiUrl)
        let articleUsecase = Usecase.ArticleUsecaseImpl(articleRepo: articleRepo)
        let authUsecase = Usecase.AuthUsecaseImpl(authRepo: authRepo)
        
        self.authUsecase = authUsecase
        self.articleUsecase = articleUsecase
    }

    var body: some Scene {
        WindowGroup {
            MainView(articleUsecase: articleUsecase, authUsecase: authUsecase)
        }
    }
}

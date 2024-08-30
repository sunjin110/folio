import GoogleSignIn
import OpenAPIRuntime
import OpenAPIURLSession
import SwiftData
import SwiftUI

@main
struct IoliooApp: App {
    let authUsecase: Usecase.AuthUsecase
    let articleUsecase: Usecase.ArticleUsecase
    let mediaUsecase: Usecase.MediaUsecase

    internal init() {
        let appConf = AppConfiguration.shared

        // default date transcoder doen't support milli sec
        let openApiConfig = OpenAPIRuntime.Configuration(
            dateTranscoder: OpenAPIRuntime.ISO8601DateTranscoder.iso8601WithFractionalSeconds)

        let serverURL = Foundation.URL(string: appConf.golioApiUrl)!
        let golioClient: APIProtocol = Client(
            serverURL: serverURL, configuration: openApiConfig, transport: URLSessionTransport())

        let articleRepo = InfraRepo.Article(client: golioClient)
        let authRepo = InfraRepo.Auth(baseUrl: appConf.golioApiUrl)
        let mediaRepo = InfraRepo.Media(client: golioClient)
        let articleUsecase = Usecase.ArticleUsecaseImpl(articleRepo: articleRepo)
        let authUsecase = Usecase.AuthUsecaseImpl(authRepo: authRepo)
        let mediaUsecase = Usecase.MediaUsecaseImpl(mediaRepo: mediaRepo)

        self.authUsecase = authUsecase
        self.articleUsecase = articleUsecase
        self.mediaUsecase = mediaUsecase
        
        print("Iolioo start")
    }

    var body: some Scene {
        WindowGroup {
            MainView(articleUsecase: articleUsecase, authUsecase: authUsecase, mediaUsecase: mediaUsecase)
        }
    }
}

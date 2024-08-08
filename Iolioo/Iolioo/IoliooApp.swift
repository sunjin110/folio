import OpenAPIURLSession
import SwiftData
import SwiftUI
import GoogleSignIn

@main
struct IoliooApp: App {
    var sharedModelContainer: ModelContainer = {
        let schema = Schema([
            Item.self
        ])
        let modelConfiguration = ModelConfiguration(schema: schema, isStoredInMemoryOnly: false)

        do {
            return try ModelContainer(for: schema, configurations: [modelConfiguration])
        } catch {
            fatalError("Could not create ModelContainer: \(error)")
        }
    }()

    var body: some Scene {
        let serverURL = Foundation.URL(string: "http://localhost:3000")!
        let golioClient: APIProtocol = Client(
            serverURL: serverURL, transport: URLSessionTransport())
        let articleRepo = InfraRepo.Article(client: golioClient)
        let authRepo = InfraRepo.Auth(baseUrl: "http://localhost:3000")

        let articleUsecase = Usecase.ArticleUsecaseImpl(articleRepo: articleRepo)
        let authUsecase = Usecase.AuthUsecaseImpl(authRepo: authRepo)

        WindowGroup {
            MainView(articleUsecase: articleUsecase, authUsecase: authUsecase)
//            .onOpenURL { url in
//                GIDSignIn.sharedInstance.handle(url)
//            }
//            .onAppear {
//                GIDSignIn.sharedInstance.restorePreviousSignIn { user, error in
//                    if error != nil {
//                        print("failed restorePreviousSignIn. err: \(error.debugDescription)")
//                    }
//                    
//                    print("==== restorePreviousSignInを通りました. user is \(user?.userID ?? "")")
//                    print("tokenid is \(user?.accessToken.tokenString ?? "nothing")")
//                }
//            }
        }
        .modelContainer(sharedModelContainer)
    }
}


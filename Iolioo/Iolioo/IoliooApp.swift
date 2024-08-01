import SwiftUI
import SwiftData
import OpenAPIURLSession

@main
struct IoliooApp: App {
    var sharedModelContainer: ModelContainer = {
        let schema = Schema([
            Item.self,
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
        let golioClient: APIProtocol = Client(serverURL: serverURL, transport: URLSessionTransport())
        let articleRepo = InfraRepo.Article(client: golioClient)
        
        let articleUsecase = Usecase.ArticleUsecaseImpl(articleRepo: articleRepo)
        
        WindowGroup {
            MainView(articleUsecase: articleUsecase)
        }
        .modelContainer(sharedModelContainer)
    }
}

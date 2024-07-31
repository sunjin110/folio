import SwiftUI
import SwiftData
import OpenAPIRuntime
import OpenAPIURLSession
import Foundation

@main
struct IolioApp: App {
    var sharedModelContainer: ModelContainer = {
        
        let serverURL = Foundation.URL(string: "http://localhost:3000")!
        
        let golioClient: APIProtocol = Client(serverURL: serverURL, transport: URLSessionTransport())
        let articleRepo = InfraRepo.Article(client: golioClient)
        let articleApplication = Application.ArticleApplicationImpl(articleRepo: articleRepo)
        
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
        WindowGroup {
            ContentView()
        }
        .modelContainer(sharedModelContainer)
    }
}

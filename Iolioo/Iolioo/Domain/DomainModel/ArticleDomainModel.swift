import Foundation

extension DomainModel {
    public struct Article {
        var id: String
        var title: String
        var body: String
        var writer: String
        var tags: [ArticleTag]
        var createdAt: Date
        var updatedAt: Date
    }

    public struct ArticleSummary {
        var id: String
        var title: String
        var tags: [ArticleTag]
        var createdAt: Date
        var updatedAt: Date
    }
}

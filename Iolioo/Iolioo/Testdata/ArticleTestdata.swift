import Foundation

extension Testdata {
    static func GetArticleSummaries() -> [DomainModel.ArticleSummary] {
        let tags = GetTags(count: 3)
        var summaries: [DomainModel.ArticleSummary] = []
        for i in 1..<10 {
            let summary = DomainModel.ArticleSummary.init(
                id: "id_\(i)", title: "title_\(i)", tags: tags, createdAt: Date.now,
                updatedAt: Date.now)

            summaries.append(summary)
        }
        return summaries
    }

    static func GetArticle() -> DomainModel.Article {
        let tags = GetTags(count: 3)
        return DomainModel.Article(
            id: "id", title: "title", body: "body", writer: "writer", tags: tags,
            createdAt: Date.now, updatedAt: Date.now)
    }

    static func GetTags(count: Int) -> [DomainModel.ArticleTag] {
        var tags: [DomainModel.ArticleTag] = []

        for i in 0..<count {
            let tag = DomainModel.ArticleTag(id: "id_\(i)", name: "tag_\(i)")
            tags.append(tag)
        }
        return tags
    }
}

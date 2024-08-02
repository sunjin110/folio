import Foundation

extension Testdata {
    static func GetArticleSummaries() -> [DomainModel.ArticleSummary] {
        var tags: [DomainModel.ArticleTag] = []
        for i in 1..<3 {
            let tag = DomainModel.ArticleTag(id: "id_\(i)", name: "tag_\(i)")
            tags.append(tag)
        }
        
        var summaries: [DomainModel.ArticleSummary] = []
        for i in 1..<10 {
            let summary = DomainModel.ArticleSummary.init(id: "id_\(i)", title: "title_\(i)", tags: tags, createdAt: Date.now, updatedAt: Date.now)
            
            summaries.append(summary)
        }
        return summaries
    }
}

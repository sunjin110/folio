import SwiftUI

struct ArticleListTemplate: View {
    
    var summaries: [DomainModel.ArticleSummary]
    
    var body: some View {
        Group {
            if summaries.isEmpty {
                Text("no articles")
            } else {
                ArticleCardsViewOrganisms(summaries: summaries)
            }
        }
        
    }
}

#Preview {
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
    return ArticleListTemplate(summaries: summaries)
}

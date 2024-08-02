import SwiftUI

struct ArticleListTemplate: View {

    var summaries: [DomainModel.ArticleSummary]
    let destinationProvider: (DomainModel.ArticleSummary) -> AnyView

    var body: some View {
        Group {
            if summaries.isEmpty {
                Text("no articles")
            } else {
                NavigationStack {
                    ArticleCardsViewOrganisms(
                        summaries: summaries, destinationProvider: destinationProvider
                    )
                    .navigationTitle("Article")
                    .toolbar {
                        Button(action: {}) {
                            Image(systemName: "plus")
                        }
                        .accessibilityLabel("New Article")
                    }
                }
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
        let summary = DomainModel.ArticleSummary.init(
            id: "id_\(i)", title: "title_\(i)", tags: tags, createdAt: Date.now, updatedAt: Date.now
        )

        summaries.append(summary)
    }

    return ArticleListTemplate(
        summaries: summaries,
        destinationProvider: { summary in
            return AnyView(Text(summary.title))
        })
}

import SwiftUI

struct ArticleCardsViewOrganisms: View {
    let summaries: [DomainModel.ArticleSummary]
    let destinationProvider: (DomainModel.ArticleSummary) -> AnyView

    var body: some View {
        List(self.summaries, id: \.id) { summary in
            NavigationLink(destination: destinationProvider(summary)) {
                ArticleCardViewMolecules(summary: summary)
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

    return ArticleCardsViewOrganisms(
        summaries: summaries,
        destinationProvider: { summary in
            return AnyView(Text(summary.title))
        })
}

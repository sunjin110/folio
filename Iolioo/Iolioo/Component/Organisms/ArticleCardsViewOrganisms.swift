import SwiftUI

struct ArticleCardsViewOrganismsLoadMoreOutput {
    let isFinished: Bool
}

struct ArticleCardsViewOrganisms: View {
    @Binding var summaries: [DomainModel.ArticleSummary]
    let destinationProvider: (DomainModel.ArticleSummary) -> AnyView
    
    let loadMoreFunc: (() async -> ArticleCardsViewOrganismsLoadMoreOutput)?

    @State private var isLoading = false
    @State private var isFinished = false
    
    var body: some View {
        List {
            ForEach(self.summaries, id: \.id) { summary in
                NavigationLink(destination: destinationProvider(summary)) {
                    ArticleCardViewMolecules(summary: summary)
                }
            }

            if loadMoreFunc != nil && !isFinished {
                ProgressView()
                    .frame(maxWidth: .infinity, maxHeight: .infinity)
                    .task {
                        guard let loadMoreFunc = loadMoreFunc else {
                            return
                        }
                        
                        if isLoading {
                            return
                        }
                        
                        isLoading = true
                        let output = await loadMoreFunc()
                        self.isFinished = output.isFinished
                        isLoading = false
                    }
            }
        }
    }
}

#if DEBUG
struct ArticleCardsViewOrganisms_Previews: PreviewProvider {
    struct PreviewWrapper: View {
        @State var summaries: [DomainModel.ArticleSummary] = []
        
        init() {
            // 初期データの設定
            let tags: [DomainModel.ArticleTag] = (1..<3).map { i in
                DomainModel.ArticleTag(id: "id_\(i)", name: "tag_\(i)")
            }
            
            _summaries = State(initialValue: (1..<10).map { i in
                DomainModel.ArticleSummary(
                    id: "id_\(i)", title: "title_\(i)", tags: tags, createdAt: Date(), updatedAt: Date()
                )
            })
        }
        
        var body: some View {
            ArticleCardsViewOrganisms(
                summaries: $summaries,
                destinationProvider: { summary in
                    AnyView(Text(summary.title))
                },
                loadMoreFunc: {
                    let additionalSummaries = (10..<20).map { i in
                        DomainModel.ArticleSummary(
                            id: "id_\(i)",
                            title: "title_\(i)",
                            tags: [],
                            createdAt: Date(),
                            updatedAt: Date()
                        )
                    }
                    summaries.append(contentsOf: additionalSummaries)
                    return .init(isFinished: true)
                }
            )
        }
    }
    
    static var previews: some View {
        PreviewWrapper()
    }
}
#endif

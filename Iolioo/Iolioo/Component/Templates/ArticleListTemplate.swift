import SwiftUI

struct ArticleListTemplateLoadMoreOutput {
    let isFinished: Bool
}

struct ArticleListTemplate: View {

    @Binding var summaries: [DomainModel.ArticleSummary]
    let destinationProvider: (DomainModel.ArticleSummary) -> AnyView
    let createArticleView: AnyView
    let loadMoreArticlesFunc: () async -> ArticleListTemplateLoadMoreOutput
    let refreshArticlesFunc: () async -> Void

    var body: some View {
        Group {
            NavigationStack {
                ArticleCardsViewOrganisms(
                    summaries: $summaries, destinationProvider: destinationProvider,
                    loadMoreFunc: {
                        let output = await loadMoreArticlesFunc()
                        return .init(isFinished: output.isFinished)
                    },
                    refreshFunc: refreshArticlesFunc
                )
                .navigationTitle("Article")
                .toolbar {
                    NavigationLink(destination: {
                        createArticleView
                    }) {
                        Image(systemName: "plus")
                    }.accessibilityLabel("New Article")
                }
            }
        }
    }
}

#if DEBUG
struct ArticleListTemplate_Previews: PreviewProvider {
    struct PreviewWrapper: View {
        @State var summaries: [DomainModel.ArticleSummary] = []
        @State var isFinished = false
        @State var isLoading = false
        
        init() {
            let tags = (1..<3).map { i in
                DomainModel.ArticleTag(id: "id_\(i)", name: "name_\(i)")
            }
            
            _summaries = State(initialValue: (1..<10).map { i in
                DomainModel.ArticleSummary(id: "id_\(i)", title: "title_\(i)", tags: tags, createdAt: Date.now, updatedAt: Date.now)
            })
        }
        
        var body: some View {
            
            ArticleListTemplate(summaries: $summaries, destinationProvider: { summary in
                return AnyView(Text(summary.title))}, createArticleView: AnyView(Text("Create Article View")), loadMoreArticlesFunc: self.loadMoreArticles, refreshArticlesFunc: refreshArticles)
        }
        
        func loadMoreArticles() async -> ArticleListTemplateLoadMoreOutput {
            let tags = (1..<3).map { i in
                DomainModel.ArticleTag(id: "id_\(i)", name: "name_\(i)")
            }
            
            let additionalSummaries = (10..<20).map { i in
                DomainModel.ArticleSummary(id: "id_\(i)", title: "title_\(i)", tags: tags, createdAt: Date.now, updatedAt: Date.now)
            }
            summaries.append(contentsOf: additionalSummaries)
            
            return .init(isFinished: false)
        }
        
        private func refreshArticles() async -> Void {
            let tags = (1..<3).map { i in
                DomainModel.ArticleTag(id: "id_\(i)", name: "name_\(i)")
            }
            
            summaries = (1..<10).map { i in
                DomainModel.ArticleSummary(id: "id_\(i)", title: "title_\(i)", tags: tags, createdAt: Date.now, updatedAt: Date.now)
            }
        }
    }
    
    static var previews: some View {
        PreviewWrapper()
    }
}
#endif

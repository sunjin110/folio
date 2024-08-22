import SwiftUI

struct ArticleListView: View {

    let articleUsecase: Usecase.ArticleUsecase
    @State var summaries: [DomainModel.ArticleSummary]
    
    let limit = 10
    
    var offset: Int {
        self.summaries.count
    }

    var body: some View {
        ArticleListTemplate(
            summaries: $summaries,
            destinationProvider: { summary in
                AnyView(ArticleDetailView(articleUsecase: articleUsecase, id: summary.id))
            },
            createArticleView: AnyView(ArticleCreateView(articleUsecase: articleUsecase)),
            loadMoreArticlesFunc: self.loadMoreArticles,
            refreshArticlesFunc: refreshArticles
        )
    }
    
    private func loadMoreArticles() async -> ArticleListTemplateLoadMoreOutput {
        switch await articleUsecase.find(offset: self.offset, limit: limit, searchTitleText: nil) {
        case .success(let summaries):
            self.summaries.append(contentsOf: summaries)
            if summaries.count < limit {
                return .init(isFinished: true)
            }
            return .init(isFinished: false)
        case .failure(let err):
            print("error: \(err)")
            return .init(isFinished: false)
        }
    }
    
    private func refreshArticles() async {
        switch await articleUsecase.find(offset: 0, limit: limit, searchTitleText: nil) {
        case .success(let summaries):
            self.summaries = summaries
            if summaries.count < limit {
                return
            }
            return
        case .failure(let err):
            print("error: \(err)")
            return
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

    var articleUsecase = Usecase.ArticleUsecaseMock()
    articleUsecase.findResult = .success(summaries)
    articleUsecase.getResult = .success(Testdata.GetArticle())

    return ArticleListView(articleUsecase: articleUsecase, summaries: [])
}

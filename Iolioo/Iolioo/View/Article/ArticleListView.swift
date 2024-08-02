import SwiftUI

struct ArticleListView: View {

    let articleUsecase: Usecase.ArticleUsecase
    @State var summaries: [DomainModel.ArticleSummary]

    var body: some View {
        ArticleListTemplate(
            summaries: summaries,
            destinationProvider: { summary in
                AnyView(ArticleDetailView(articleUsecase: articleUsecase, id: summary.id))
            }
        ).task {
            let result = await articleUsecase.find(offset: 0, limit: 10, searchTitleText: nil)
            switch result {
            case .success(let summaries):
                self.summaries = summaries
            case .failure(let err):
                print("error: \(err)")
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

    var articleUsecase = Usecase.ArticleUsecaseMock()
    articleUsecase.findResult = .success(summaries)
    articleUsecase.getResult = .success(Testdata.GetArticle())

    return ArticleListView(articleUsecase: articleUsecase, summaries: [])
}

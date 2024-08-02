import SwiftUI

struct ArticleDetailView: View {

    var articleUsecase: Usecase.ArticleUsecase
    var id: String
    @State var article: DomainModel.Article?

    var body: some View {

        ArticleDetailTemplate(article: self.article).task {
            await loadArticleDetail()
        }
    }

    private func loadArticleDetail() async {
        let result = await self.articleUsecase.get(id: id)
        switch result {
        case .success(let article):
            self.article = article
        case .failure(_):
            self.article = nil
        }
    }
}

#Preview {
    let articleUsecase = Usecase.ArticleUsecaseMock(
        getResult: .success(
            DomainModel.Article(
                id: "id", title: "title", body: "body", writer: "writer", tags: [],
                createdAt: Date.now, updatedAt: Date.now)))
    return ArticleDetailView(articleUsecase: articleUsecase, id: "id")
}

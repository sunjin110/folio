import SwiftUI

struct ArticleDetailView: View {

    let articleUsecase: Usecase.ArticleUsecase
    var id: String
    @State var article: DomainModel.Article?

    var body: some View {
        ArticleDetailTemplate(article: $article, editDestinationFunc: self.editDestination).task {
            await loadArticleDetail()
        }
    }

    private func loadArticleDetail() async {
        let result = await self.articleUsecase.get(id: id)
        switch result {
        case .success(let article):
            self.article = article
        case .failure(let err):
            print("fialed self.articleUsecase.get(id: id). id: \(id), err: \(err)")
            self.article = nil
        }
    }
    
    private func editDestination(article: DomainModel.Article) -> AnyView {
        AnyView(ArticleUpdateView(articleUsecase: self.articleUsecase, id: article.id))
    }
}

#Preview {
    let articleUsecase = Usecase.ArticleUsecaseMock(
        getResult: .success(
            DomainModel.Article(
                id: "id", title: "title", body: "body", writer: "writer", tags: [],
                createdAt: Date.now, updatedAt: Date.now)))
    return NavigationStack {
        ArticleDetailView(articleUsecase: articleUsecase, id: "id")
    }
}

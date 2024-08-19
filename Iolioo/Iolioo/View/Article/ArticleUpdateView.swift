import SwiftUI
import AlertToast

struct ArticleUpdateView: View {
    let articleUsecase: Usecase.ArticleUsecase
    let id: String
    
    @State
    private var showSaveSuccessToast = false
    
    @State
    private var showSaveFailedToast = false
    
    @State
    private var article: DomainModel.Article?
    
    var body: some View {
        ArticleUpdateTemplate(article: $article, saveArticleFunc: self.saveArticle).task {
            await self.loadArticleDetail()
        }
        .toast(isPresenting: $showSaveSuccessToast, alert: {
            AlertToast(displayMode: .hud, type: .complete(.green), title: "Save Success!")
        })
        .toast(isPresenting: $showSaveFailedToast, alert: {
            AlertToast(displayMode: .hud, type: .error(.red), title: "Failed save article")
        })
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
    
    private func saveArticle(input: (title: String, body: String)) async -> Void {
        guard var article = self.article else {return}
        
        article.title = input.title
        article.body = input.body
        
        let result = await self.articleUsecase.update(article: article)
        switch result {
        case .success(_):
            showSaveSuccessToast = true
        case .failure(let err):
            print("failed self.articleUsecase.update. article: \(article), err: \(err)")
            showSaveFailedToast = true
        }
    }
}

#Preview {
    
    var articleUsecase = Usecase.ArticleUsecaseMock()
    
    articleUsecase.getResult = .success(DomainModel.Article(id: "id", title: "title", body: "本文です\n本文みたいです", writer: "writer", tags: [], createdAt: Date.now, updatedAt: Date.now))
    
    
    return NavigationStack {
        ArticleUpdateView(articleUsecase: articleUsecase, id: "id")
    }
}

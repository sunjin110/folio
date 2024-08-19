import SwiftUI

struct ArticleCreateView: View {
    
    let articleUsecase: Usecase.ArticleUsecase
    
    var body: some View {
        ArticleCreateTemplate(saveArticleFunc: saveArticle)
    }
    
    private func saveArticle(input: (title: String, body: String)) async -> Void {
        let result = await self.articleUsecase.insert(title: input.title, body: input.body)
        
        switch result {
        case .success(_):
            return
        case .failure(let err):
            print("fialed save article. err: \(err)")
            return
        }
    }
}

#Preview {
    let articleUsecase = Usecase.ArticleUsecaseMock()
    return ArticleCreateView(articleUsecase: articleUsecase)
}

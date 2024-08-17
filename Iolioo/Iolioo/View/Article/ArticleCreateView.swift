import SwiftUI

struct ArticleCreateView: View {
    
    let articleUsecase: Usecase.ArticleUsecase
    
    var body: some View {
        ArticleCreateTemplate(saveArticleFunc: saveArticle)
    }
    
    private func saveArticle(input: (title: String, body: String)) async -> Void {
        print("todo")
    }
}

#Preview {
    var articleUsecase = Usecase.ArticleUsecaseMock()
    return ArticleCreateView(articleUsecase: articleUsecase)
}

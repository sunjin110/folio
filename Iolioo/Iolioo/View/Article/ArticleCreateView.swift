import SwiftUI
import AlertToast

struct ArticleCreateView: View {
    
    let articleUsecase: Usecase.ArticleUsecase
    
    @State
    private var showSaveSuccessToast = false
    
    @State
    private var showSaveFailedToast = false
    
    var body: some View {
        ArticleCreateTemplate(saveArticleFunc: saveArticle)
            .toast(isPresenting: $showSaveSuccessToast, alert: {
                AlertToast(displayMode: .hud, type: .complete(.green), title: "Save Success!")
            })
            .toast(isPresenting: $showSaveFailedToast, alert: {
                AlertToast(displayMode: .hud, type: .error(.red), title: "Failed save article")
            })
    }
    
    private func saveArticle(input: (title: String, body: String)) async -> Void {
        let result = await self.articleUsecase.insert(title: input.title, body: input.body)
        
        switch result {
        case .success(_):
            showSaveSuccessToast = true
            return
        case .failure(let err):
            print("fialed save article. err: \(err)")
            showSaveFailedToast = true
            return
        }
    }
}

#Preview {
    let articleUsecase = Usecase.ArticleUsecaseMock()
    return ArticleCreateView(articleUsecase: articleUsecase)
}

import SwiftUI

struct ArticleUpdateTemplate: View {
    
    @Binding
    var article: DomainModel.Article?
    
    @State
    private var title: String = ""
    
    @State
    private var articleBody: String = ""
    
    @State
    private var articleWasLoaded = false
    
    @State
    private var isSaving = false
    
    let saveArticleFunc: ((title: String, body: String)) async -> Void
    
    var body: some View {
        Group {
            if self.article == nil {
                Text("not found article")
            } else {
                ScrollView {
                    VStack(alignment: .leading, content: {
                        TextField("Title", text: $title).font(.largeTitle).bold().padding(.bottom)
                        TextField("Body", text: $articleBody, axis: .vertical)
                    })
                    .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .leading)
                }
            }
        }
        .padding(.horizontal)
        .toolbar {
            ToolbarItem(placement: .topBarTrailing) {
                Button(action: {
                    self.saveArticle()
                }) {
                    Text("Save")
                }
            }
        }
        .onReceive([self.article].publisher.first()) {data in
            if self.articleWasLoaded {
                return
            }
            
            if let article = data {
                self.title = article.title
                self.articleBody = article.body
                self.articleWasLoaded = true
            }
        }
    }
    
    private func saveArticle() {
        guard !isSaving else { return }
        
        isSaving = true
        Task {
            await saveArticleFunc((title: self.title, body: self.articleBody))
            isSaving = false
        }
    }
}

#if DEBUG
#Preview {
    
    func saveArticlePreview(input: (title: String, body: String)) async -> Void {
        print("save! title is \(input.title), body is \(input.body)")
    }
    
    @State var article: DomainModel.Article? = DomainModel.Article(id: "id", title: "title", body: "body", writer: "wirter", tags: [], createdAt: Date.now, updatedAt: Date.now)
    return NavigationStack {
        ArticleUpdateTemplate(article: $article, saveArticleFunc: saveArticlePreview)
    }
}
#endif

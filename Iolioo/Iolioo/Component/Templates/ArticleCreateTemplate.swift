import SwiftUI

struct ArticleCreateTemplate: View {
    
    @State
    private var title: String = ""
    
    @State
    private var articleBody: String = ""
    
    @State
    private var tapSaveButton: Bool = false
    
    // 2重タップ防止
    @State
    private var isSavingArticle: Bool = false
    
    let saveArticleFunc: ((title: String, body: String)) async -> Void
    
    var body: some View {
        Group {
            ScrollView {
                VStack(
                    alignment: .leading, content: {
                        TextField("😎Title", text: $title).font(.largeTitle).bold()
                            .padding(.bottom)
                        TextField("🕺Body", text: $articleBody, axis: .vertical)
                    }
                )
                    .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .leading)
            }
        }
        .padding(.horizontal)
        .toolbar {
            ToolbarItem(placement: .navigationBarTrailing) {
                
                Button(action: {
                    self.tapSaveButton = true
                }) {
                    Text("Save")
                }
            }
        }
        .task(id: self.tapSaveButton) {
            if self.isSavingArticle {
                return
            }
            isSavingArticle = true
            await self.saveArticleFunc((title: self.title, body: self.articleBody))
            isSavingArticle = false
        }
    }
}

#if DEBUG
#Preview {
    func saveArticlePreview(input: (title: String, body: String)) async -> Void {
        print("save! title is \(input.title), body is \(input.body)")
    }
    return ArticleCreateTemplate(saveArticleFunc: saveArticlePreview)
}
#endif
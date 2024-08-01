import SwiftUI

struct ArticleDetailTemplateComponent: View {
    var article: DomainModel.Article?
    var body: some View {
        Group {
            if self.article == nil {
                Text("not found article")
            } else {
                VStack(alignment: .leading, content: {
                    Text("title: \(self.article!.title)")
                    Text("body: \(self.article!.body)")
                    Text("created_at: \(self.article!.createdAt)")
                })
            }
        }
    }
}

#Preview {
    let article = DomainModel.Article(id: "id", title: "title", body: "body", writer: "writer", tags: [], createdAt: Date.now, updatedAt: Date.now)
    return ArticleDetailTemplateComponent(article: article)
}

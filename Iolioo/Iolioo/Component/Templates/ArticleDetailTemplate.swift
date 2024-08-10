import MarkdownView
import SwiftUI

struct ArticleDetailTemplate: View {
    @Binding var article: DomainModel.Article?
    var body: some View {
        Group {
            if self.article == nil {
                Text("not found article")
            } else {
                ScrollView {
                    VStack(
                        alignment: .leading,
                        content: {
                            Text(self.article!.title).font(.largeTitle).bold().padding(
                                .bottom)
                            MarkdownView(text: self.article!.body)
                            Spacer()
                        }
                    ).padding(.horizontal)
                        .frame(maxWidth: .infinity, maxHeight: .infinity, alignment: .leading)
                }
            }
        }.toolbar {
            ToolbarItem(placement: .navigationBarTrailing) {
                Button(action: {
                    print("pressed edit button")
                }) {
                    Text("Edit")
                }
            }
            ToolbarItem {
                Button(action: {
                    print("plessed add button")
                }) {
                    Label("Add Item", systemImage: "plus")
                }
            }
        }
    }
}

#Preview {

    let body = """
        # body title
        ## body title2

        ### task
        - [ ] task1
        - [x] task2

        > Quote and `inline code`

        This is the Apple's **newly published** [swift-markdown](https://github.com/apple/swift-markdown)
        """

    @State var article: DomainModel.Article? = DomainModel.Article(
        id: "id", title: "title", body: body, writer: "writer",
        tags: [], createdAt: Date.now,
        updatedAt: Date.now)
    return NavigationStack {
        ArticleDetailTemplate(article: $article)
    }
}

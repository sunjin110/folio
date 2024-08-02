import SwiftUI
import Foundation

// ref
// https://developer.apple.com/tutorials/app-dev-training/creating-a-card-view
struct ArticleCardViewMolecules: View {
    let id = UUID()
    let summary: DomainModel.ArticleSummary
    
    var body: some View {
        VStack(alignment: .leading, content: {
            Text(summary.title).font(.headline)
            Spacer()
            HStack(content: {
                FlowLayout(alignment: .leading, spacing: 3) {
                    ForEach(self.summary.tags, id: \.id) { tag in
                        Text(tag.name)
                            .padding(.vertical, 2)
                            .padding(.horizontal, 5)
                            .background(Color(.systemGroupedBackground))
                            .cornerRadius(15)
                    }
                }
                Spacer()
                Label(self.summary.createdAt.formatted(), systemImage: "clock")
                
            })
            .font(.caption)
        })
        .padding()
    }
}

struct ArticleCardViewMolecules_Previews: PreviewProvider {
  static var previews: some View {
      
      let tags: [DomainModel.ArticleTag] = [DomainModel.ArticleTag(id: "tag_id", name: "tag_name"), DomainModel.ArticleTag(id: "tag_id", name: "tag_name")]
      
      ArticleCardViewMolecules(summary: DomainModel.ArticleSummary.init(id: "id", title: "📓 English Diary: Yesterday's Happen", tags: tags, createdAt: Date.now, updatedAt: Date.now)).background().previewLayout(.fixed(width: 400, height: 60))
  }
}

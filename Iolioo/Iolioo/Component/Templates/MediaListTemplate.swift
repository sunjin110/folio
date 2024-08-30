import SwiftUI

struct MediaListTemplate: View {
    @Binding var summaries: [DomainModel.MediumSummary]
    let loadMoreMediaFunc: () async -> Bool
    let refreshMediaFunc: () async -> Void
    
    var body: some View {
        Group {
            NavigationStack {
                MediaImageGridViewOrganisms(
                        summaries: $summaries,
                        loadMoreFunc: loadMoreMediaFunc,
                        refreshFunc: refreshMediaFunc
                )
                    .navigationTitle("Media")
                    .toolbar{
                        Image(systemName: "arrow.triangle.2.circlepath")
                    }
            }
        }
    }
}

#if DEBUG
struct MediaListTemplate_Previews: PreviewProvider {
    struct PreviewWrapper: View {
        @State var summaries: [DomainModel.MediumSummary] = []
        
        init() {
            _summaries = State(initialValue: (1..<40).map { i in
                DomainModel.MediumSummary(id: "id_\(i)", fileType: "file_type", thumbnailUrl: "https://cdn.pixabay.com/photo/2023/09/04/17/48/flamingos-8233303_1280.jpg", createdAt: Date.now, updatedAt: Date.now)
            })
        }
        
        var body: some View {
            MediaListTemplate(summaries: $summaries, loadMoreMediaFunc: loadMoreMedia, refreshMediaFunc: self.refreshMedia)
        }
        
        private func refreshMedia() {
            print("refresh!")
        }
        
        private func loadMoreMedia() -> Bool {
            return true
        }
    }
    
    static var previews: some View {
        PreviewWrapper()
    }
}

#endif

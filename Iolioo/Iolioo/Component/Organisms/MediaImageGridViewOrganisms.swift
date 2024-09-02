import SwiftUI

struct MediaImageGridViewOrganisms: View {
        
    @Binding var summaries: [DomainModel.MediumSummary]
    
    let loadMoreFunc: (() async -> Bool)?
    let refreshFunc: (() async -> Void)
    let detailDestinationProvider: (DomainModel.MediumSummary) -> AnyView

    
    @State private var isLoading = false
    @State private var isFinished = false
    
    var body: some View {
        ScrollView {
            LazyVGrid(
                columns: Array(repeating: .init(.flexible(minimum: 50, maximum: 200)), count: 4), spacing: 8, content: {
                    ForEach(self.summaries, id: \.id) { summary in
                        if let thumbnailUrl = summary.thumbnailUrl {
                            NavigationLink(destination: {
                                detailDestinationProvider(summary)
                            }) {
                                MediaImageViewMolecules(url: URL(string: thumbnailUrl))
                            }
                        }
                    }
                    
                    if loadMoreFunc != nil && !isFinished {
                        ProgressView()
                            .frame(maxWidth: .infinity, maxHeight: .infinity)
                            .task {
                                guard let loadMoreFunc = loadMoreFunc else {
                                    return
                                }
                                
                                if isLoading {
                                    return
                                }
                                
                                isLoading = true
                                self.isFinished = await loadMoreFunc()
                                isLoading = false
                            }
                    }
                    
                }
            )
        }.refreshable {
            await self.refreshFunc()
        }
    }
}

#if DEBUG
struct MediaImageGridViewOrganisms_Previews: PreviewProvider {
    struct PreviewWrapper: View {
        
        @State var summaries: [DomainModel.MediumSummary] = []
        
        init() {
            _summaries = State(initialValue: (1..<40).map { i in
                DomainModel.MediumSummary(id: "id_\(i)", fileType: "file_type", thumbnailUrl: "https://cdn.pixabay.com/photo/2023/09/04/17/48/flamingos-8233303_1280.jpg", createdAt: Date.now, updatedAt: Date.now)
            })
        }
        
        private func destinationProvider(medium: DomainModel.MediumSummary) -> AnyView {
            AnyView(Text("todo"))
        }
        
        var body: some View {
            MediaImageGridViewOrganisms(summaries: $summaries, loadMoreFunc: nil, refreshFunc: refresh, detailDestinationProvider: destinationProvider)
        }
        
        private func refresh() {
            print("refresh")
        }
    }
    
    static var previews: some View {
        NavigationStack {
            PreviewWrapper()
        }
    }
}

#endif

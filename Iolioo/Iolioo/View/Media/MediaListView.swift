import SwiftUI

struct MediaListView: View {
    
    let mediaUsecase: Usecase.MediaUsecase
    @State var summaries: [DomainModel.MediumSummary] = []
    
    let limit = 20
    
    var offset: Int {
        self.summaries.count
    }
    
    var body: some View {
        MediaListTemplate(summaries: $summaries, loadMoreMediaFunc: self.loadMoreMedia, refreshMediaFunc: self.refreshMedia, detailDestinationProvider: self.detailDestinationProvider)
    }
    
    private func loadMoreMedia() async -> Bool {
        switch await mediaUsecase.find(offset: self.offset, limit: self.limit) {
        case .success(let summaries):
            self.summaries.append(contentsOf: summaries)
            return summaries.count < limit
        case .failure(let err):
            print("error: \(err)")
            return false
        }
    }
    
    private func refreshMedia() async {
        switch await mediaUsecase.find(offset: 0, limit: limit) {
        case .success(let summaries):
            self.summaries = summaries
        case .failure(let err):
            print("error: \(err)")
        }
    }
    
    private func detailDestinationProvider(medium: DomainModel.MediumSummary) -> AnyView {
        AnyView(Text("todo"))
    }
}

#Preview {
    var mediaUsecase = Usecase.MediaUsecaesMock()
    
    let summaries = (1..<10).map { i in
        DomainModel.MediumSummary(id: "id_\(i)", fileType: "file_type", thumbnailUrl: "https://cdn.pixabay.com/photo/2023/09/04/17/48/flamingos-8233303_1280.jpg", createdAt: Date.now, updatedAt: Date.now)
    }
    
    mediaUsecase.findResult = .success(summaries)
    return MediaListView(mediaUsecase: mediaUsecase)
}

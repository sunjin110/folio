import SwiftUI

struct MediaDetailView: View {
    
    let mediaUsecase: Usecase.MediaUsecase
    let id: String
    
    @State private var medium: DomainModel.Medium?
    
    var body: some View {
        MediaDetailTemplate(imageURL: "")
    }
    
    private func loadMedium() async {
        switch await self.mediaUsecase.get(id: id) {
        case .success(let medium):
            self.medium = medium
        case .failure(let err):
            print("failed self.mediaUsecase.get. id: \(id), err: \(err)")
            self.medium = nil
        }
        
    }
}

#Preview {
    let mediaUsecase = Usecase.MediaUsecaesMock()
    return MediaDetailView(mediaUsecase: mediaUsecase, id: "id")
}

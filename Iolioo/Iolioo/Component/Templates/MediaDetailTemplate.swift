import SwiftUI

struct MediaDetailTemplate: View {
    
    let imageURL: String
    
    var body: some View {
        Group {
            // 画像以外もいけるように今後する
            AsyncImage(url: URL(string: self.imageURL)) { image in
                image.resizable().scaledToFit()
            } placeholder: {
                ProgressView()
            }
            .frame(maxWidth: .infinity, maxHeight: .infinity)
        }
    }
}

#Preview {
    MediaDetailTemplate(imageURL: "https://cdn.pixabay.com/photo/2023/09/04/17/48/flamingos-8233303_1280.jpg")
}

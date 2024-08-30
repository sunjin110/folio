import SwiftUI

struct MediaImageViewMolecules: View {
    var url: URL?
    var body: some View {
        Color.clear.overlay(
                AsyncImage(url: self.url) { image in
                    image.resizable().scaledToFill()
                } placeholder: {
                    ProgressView()
                }
        )
       .frame(maxWidth: .infinity)
       .aspectRatio(1, contentMode: .fit)
       .clipped()
    }
}
    
#Preview {
    LazyVGrid(
        columns: Array(repeating: .init(.flexible(minimum: 50, maximum: 200)), count: 4), spacing: 8, content: {
            MediaImageViewMolecules(url: URL(string: "https://cdn.pixabay.com/photo/2023/09/04/17/48/flamingos-8233303_1280.jpg"))
            
            MediaImageViewMolecules(url: URL(string: "https://cdn.pixabay.com/photo/2023/09/04/17/48/flamingos-8233303_1280.jpg"))
            
            MediaImageViewMolecules(url: URL(string: "https://cdn.pixabay.com/photo/2023/09/04/17/48/flamingos-8233303_1280.jpg"))
            
            MediaImageViewMolecules()
            
            MediaImageViewMolecules(url: URL(string: "https://cdn.pixabay.com/photo/2023/09/04/17/48/flamingos-8233303_1280.jpg"))
            
            MediaImageViewMolecules(url: URL(string: "https://cdn.pixabay.com/photo/2023/09/04/17/48/flamingos-8233303_1280.jpg"))
            
            MediaImageViewMolecules(url: URL(string: "https://cdn.pixabay.com/photo/2023/09/04/17/48/flamingos-8233303_1280.jpg"))
            
            MediaImageViewMolecules(url: URL(string: "https://cdn.pixabay.com/photo/2023/09/04/17/48/flamingos-8233303_1280.jpg"))
        }
    )
}


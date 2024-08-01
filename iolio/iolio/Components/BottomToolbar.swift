import SwiftUI

struct BottomToolbar: ToolbarContent {
    var body: some ToolbarContent {
        ToolbarItem(placement: .bottomBar) {
          HStack {
              Spacer()
              Button("Home", action: {
                  // TODO
              })
              Spacer()
              Button("Article", action: {
                  // TODO
              })
              Spacer()
              Button("Media", action: {
                  // TODO
              })
              Spacer()
              Button("Setting", action: {
                  // TODO
              })
              Spacer()
          }
      }
    }
}

#Preview {
    
    struct PreviewBottomToolbar: View {
        var body: some View {
            NavigationView {
                Text("test").toolbar{
                    BottomToolbar()
                }
            }
        }
    }
    
    return PreviewBottomToolbar()
}

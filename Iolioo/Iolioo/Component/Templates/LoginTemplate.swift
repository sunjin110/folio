import SwiftUI

struct LoginTemplate: View {
    var googleSignInOnTap: () -> Void
    
    var body: some View {
        VStack {
            Spacer()
            Text("Folio").font(.largeTitle).bold()
            Text("🔐 Login 🔐").font(.title).bold()
            Spacer()
            Button(action: googleSignInOnTap, label: {
                Text("Google SignIn")
            })
            Spacer()
        }.frame(maxWidth: .infinity, maxHeight: .infinity)
//        .ignoresSafeArea(.all)
    }
}

#Preview {
    LoginTemplate(googleSignInOnTap: {})
}

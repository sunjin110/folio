import SwiftUI

struct SettingTemplate: View {

    var loginOnTap: () -> Void

    var body: some View {
        VStack {
            Spacer()
            Text("⚙️ Setting!").font(.largeTitle).bold()
            Spacer()
            Button(
                action: loginOnTap,
                label: {
                    Text("Login").font(.title)
                })
            Spacer()
        }
    }
}

#Preview {
    SettingTemplate(loginOnTap: {})
}

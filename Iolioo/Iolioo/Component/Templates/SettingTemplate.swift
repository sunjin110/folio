import SwiftUI
import AlertToast

struct SettingTemplate: View {

    var loginOnTap: () -> Void
    
    @State
    private var showTast = false
    
    @State
    private var signInShowTastFlag = false

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
            
            Button("Show Toast") {
                showTast = true
            }
            Spacer()
            
            Button("Sign in Success toast") {
                signInShowTastFlag = true
            }
        }.toast(isPresenting: $showTast) {
            AlertToast(displayMode: .hud, type: .complete(.green), title: "Hello Show Tast", subTitle: "It's made by library")
        }.toast(isPresenting: $signInShowTastFlag) {
            AlertToast(displayMode: .alert, type: .complete(.green), title: "Sign in!")
        }
    }
}

#Preview {
    SettingTemplate(loginOnTap: {})
}

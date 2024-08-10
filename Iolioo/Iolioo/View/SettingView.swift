import SwiftUI

struct SettingView: View {
    var authUsecase: Usecase.AuthUsecase
    @Binding var showLogin: Bool

    var body: some View {
        SettingTemplate(loginOnTap: self.loginOnTap)
    }

    private func loginOnTap() {
        withAnimation {
            showLogin.toggle()
        }
    }
}

#Preview {
    let authUsecase = Usecase.AuthUsecaseMock()
    @State var showLogin: Bool = false
    return SettingView(authUsecase: authUsecase, showLogin: $showLogin)
}

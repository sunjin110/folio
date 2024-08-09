import SwiftUI
import GoogleSignIn

struct LoginView: View {
    
    let authUsecase: Usecase.AuthUsecase
    
    @Binding var showLogin: Bool
    
    var body: some View {
        LoginTemplate(googleSignInOnTap: handleGoogleSignInOnTap)
    }
    
    func handleGoogleSignInOnTap() {
        guard let windowScene = UIApplication.shared.connectedScenes.first as? UIWindowScene,
              let rootViewController = windowScene.windows.first(where: { $0.isKeyWindow })?.rootViewController else {
               print("RootViewController not found.")
               return
        }

        GIDSignIn.sharedInstance.signIn(withPresenting: rootViewController as UIViewController) { signInResult, error in
            if error != nil {
                print("failed google sign in. err: \(error.debugDescription)")
                return
            }
            
            guard let signInResult = signInResult else {
                print("signInResult is empty")
                return
            }
            
            signInResult.user.refreshTokensIfNeeded { user, error in
                if error != nil {
                    print("fialed signInResult.user.refreshTokensIfNeeded. error: \(error.debugDescription)")
                    return
                }
                
                guard let user = user else { print ("user is empty"); return }
                guard let idToken = user.idToken else {
                    print("id token is empty")
                    return
                }
                
                Task {
                    let result = await self.authUsecase.verifyTokenAndStartSession(idToken: idToken.tokenString, accessToken: user.accessToken.tokenString, refreshToken:user.refreshToken.tokenString)
                    switch result {
                    case .success(_):
                        DispatchQueue.main.async {
                            print("success verify and start session")
                            self.showLogin = false
                        }
                    case .failure(let err):
                        print("failed authUsecase.verifyTokenAndStartSession. err: \(err)")
                    }
                }
            }
        }
    }
    
    func handleGoogleSignOutOnTap() {
        print("pressed sign out")
        GIDSignIn.sharedInstance.signOut()
    }
}

#Preview {
    let authUsecase = Usecase.AuthUsecaseMock()
    @State var showLogin: Bool = false
    return LoginView(authUsecase: authUsecase, showLogin: $showLogin)
}

import SwiftUI
import GoogleSignIn
import GoogleSignInSwift

struct HomeView: View {
    
    var articleUsecase: Usecase.ArticleUsecase
    
    var body: some View {
        VStack {
            Text("Home").font(.title)
            Button("Google Signin", action: handleSignInButton)
            Button("Google SignOut", action: handleSignOutButton)
        }
    }
    
    func handleSignInButton() {
        guard let windowScene = UIApplication.shared.connectedScenes.first as? UIWindowScene,
                 let rootViewController = windowScene.windows.first(where: { $0.isKeyWindow })?.rootViewController else {
               print("RootViewController not found.")
               return
           }
        
        GIDSignIn.sharedInstance.signIn(withPresenting: rootViewController) { signInResult, error in
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
                
                print("id token is \(user.idToken!.tokenString)")
                print("access token is \(user.accessToken.tokenString)")
                print("refresh token is \(user.refreshToken.tokenString)")
                
                // TODO golioにid tokenを検証するやつを作る
                // https://developers.google.com/identity/sign-in/ios/backend-auth?hl=ja
                
            }
            
        }
    }
    
    func tokenSignIn(idToken: String) {
        // TODO golioに通信するところ、これはrepository層だね
    }
    
    func handleSignOutButton() {
        print("pressed sign out")
        GIDSignIn.sharedInstance.signOut()
    }
}

#Preview {
    let articleUsecase = Usecase.ArticleUsecaseMock()
    return HomeView(articleUsecase: articleUsecase)
}

import SwiftUI
import GoogleSignIn

struct MainView: View {

    @State
    var tabSelection = 1
    
    @State
    private var showLogin = false

    var articleUsecase: Usecase.ArticleUsecase
    var authUsecase: Usecase.AuthUsecase

    var body: some View {
        TabView(selection: $tabSelection) {
            HomeView(articleUsecase: articleUsecase).tabItem {
                Label("Home", systemImage: "house")
            }.tag(1)
            ArticleListView(articleUsecase: articleUsecase, summaries: []).tabItem {
                Label("Article", systemImage: "note")
            }.tag(2)
            MediaView().tabItem {
                Label("Media", systemImage: "photo.stack.fill")
            }.tag(3)
            SettingView(authUsecase: authUsecase, showLogin: $showLogin).tabItem {
                Label("Setting", systemImage: "gear")
            }.tag(4)
        }.overlay(
            Group {
                if showLogin {
                    LoginView(authUsecase: authUsecase, showLogin: $showLogin).background(.white)
                }
            }
        )
        // GoogleOAuth2.0
        .onOpenURL { url in
            GIDSignIn.sharedInstance.handle(url)
        }
        .onAppear {
            GIDSignIn.sharedInstance.restorePreviousSignIn { user, error in
                if error != nil {
                    print("failed restorePreviousSignIn. err: \(error.debugDescription)")
                    showLogin.toggle()
                    return
                }
                
                guard let user = user else {
                    showLogin.toggle()
                    return
                }
                print("==== restorePreviousSignInを通りました. user is \(user.userID ?? "user id not found")")
                print("tokenid is \(user.accessToken.tokenString)")
            }
        }
    }
}

#Preview {
    var articleUsecase = Usecase.ArticleUsecaseMock(
        getResult: .success(
            DomainModel.Article(
                id: "id", title: "title", body: "body", writer: "writer", tags: [],
                createdAt: Date.now, updatedAt: Date.now)))
    articleUsecase.findResult = .success(Testdata.GetArticleSummaries())
    
    let authUsecase = Usecase.AuthUsecaseMock()
    return MainView(articleUsecase: articleUsecase, authUsecase: authUsecase)
}

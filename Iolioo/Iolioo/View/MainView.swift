import GoogleSignIn
import SwiftUI

struct MainView: View {

    @State
    var tabSelection = 1

    @Environment(\.colorScheme) var colorSchema

    var backgroundColor: Color {
        colorSchema == .dark ? .black : .white
    }

    @State
    private var showLogin = false

    var articleUsecase: Usecase.ArticleUsecase
    var authUsecase: Usecase.AuthUsecase
    let mediaUsecase: Usecase.MediaUsecase

    var body: some View {
        TabView(selection: $tabSelection) {
            HomeView(articleUsecase: articleUsecase).tabItem {
                Label("Home", systemImage: "house")
            }.tag(1)
            ArticleListView(articleUsecase: articleUsecase, summaries: []).tabItem {
                Label("Article", systemImage: "note")
            }.tag(2)
            MediaListView(mediaUsecase: mediaUsecase).tabItem {
                Label("Media", systemImage: "photo.stack.fill")
            }.tag(3)
            SettingView(authUsecase: authUsecase, showLogin: $showLogin).tabItem {
                Label("Setting", systemImage: "gear")
            }.tag(4)
        }.overlay(
            Group {
                if showLogin {
                    LoginView(authUsecase: authUsecase, showLogin: $showLogin).background(
                        backgroundColor)
                }
            }
        )
        // GoogleOAuth2.0
        .onOpenURL { url in
            GIDSignIn.sharedInstance.handle(url)
        }
        .onAppear {
            // 起動時ログイン処理
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

                guard let idToken = user.idToken else {
                    showLogin.toggle()
                    return
                }

                Task {
                    let result = await self.authUsecase.verifyTokenAndStartSession(
                        idToken: idToken.tokenString, accessToken: user.accessToken.tokenString,
                        refreshToken: user.refreshToken.tokenString)

                    switch result {
                    case .success(_):
                        return
                    case .failure(let err):
                        print("failed authUsecase.verifyTokenAndStartSession. err: \(err)")
                    }
                }
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
    let mediaUsecase = Usecase.MediaUsecaesMock()
    return MainView(articleUsecase: articleUsecase, authUsecase: authUsecase, mediaUsecase: mediaUsecase).preferredColorScheme(
        .dark)
}

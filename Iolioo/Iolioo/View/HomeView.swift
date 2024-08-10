import GoogleSignIn
import GoogleSignInSwift
import SwiftUI

struct HomeView: View {

    var articleUsecase: Usecase.ArticleUsecase

    var body: some View {
        VStack {
            Text("Home").font(.title)
            Text("Foliooooo")
        }
    }
}

#Preview {
    let articleUsecase = Usecase.ArticleUsecaseMock()
    return HomeView(articleUsecase: articleUsecase)
}

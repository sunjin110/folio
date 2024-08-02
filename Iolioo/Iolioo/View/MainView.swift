import SwiftUI

struct MainView: View {

    @State
    var tabSelection = 1

    var articleUsecase: Usecase.ArticleUsecase

    var body: some View {
        TabView(selection: $tabSelection) {
            HomeView().tabItem {
                Label("Home", systemImage: "house")
            }.tag(1)
            ArticleListView(articleUsecase: articleUsecase, summaries: []).tabItem {
                Label("Article", systemImage: "note")
            }.tag(2)
            MediaView().tabItem {
                Label("Media", systemImage: "photo.stack.fill")
            }.tag(3)
            SettingView().tabItem {
                Label("Setting", systemImage: "gear")
            }.tag(4)
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
    return MainView(articleUsecase: articleUsecase)
}

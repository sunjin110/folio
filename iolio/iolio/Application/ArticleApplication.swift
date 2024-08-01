extension Application {
    protocol ArticleApplication {
        func get(id: String) async -> Result<DomainModel.Article, ApplicationError>
    }
}

public extension Application {
    class ArticleApplicationImpl: ArticleApplication {
        var articleRepo: DomainRepo.Article
        
        init(articleRepo: DomainRepo.Article) {
            self.articleRepo = articleRepo
        }
        
        func get(id: String) async -> Result<DomainModel.Article, Application.ApplicationError> {
            let artile = await self.articleRepo.get(id: id)
            switch artile {
            case .success(let article):
                return .success(article)
            case .failure(let err):
                return .failure(.init(message: "failed get article. id: \(id)", innerError: err, kind: .internalError))
            }
        }
    }
}

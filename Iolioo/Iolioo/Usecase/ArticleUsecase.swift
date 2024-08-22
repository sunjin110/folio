extension Usecase {
    protocol ArticleUsecase {
        func get(id: String) async -> Result<DomainModel.Article, UsecaseError>
        func find(offset: Int, limit: Int, searchTitleText: String?) async -> Result<
            [DomainModel.ArticleSummary], UsecaseError
        >
        func insert(title: String, body: String) async -> Result<(), UsecaseError>
        func update(article: DomainModel.Article) async -> Result<(), UsecaseError>
    }
}

extension Usecase {
    class ArticleUsecaseImpl: ArticleUsecase {

        var articleRepo: DomainRepo.Article

        init(articleRepo: DomainRepo.Article) {
            self.articleRepo = articleRepo
        }

        func get(id: String) async -> Result<DomainModel.Article, UsecaseError> {
            let artile = await self.articleRepo.get(id: id)
            switch artile {
            case .success(let article):
                return .success(article)
            case .failure(let err):
                return .failure(
                    .init(
                        message: "failed get article. id: \(id)", innerError: err,
                        kind: .internalError))
            }
        }

        func find(offset: Int, limit: Int, searchTitleText: String?) async -> Result<
            [DomainModel.ArticleSummary], Usecase.UsecaseError
        > {
            let articles = await self.articleRepo.find(
                offset: offset, limit: limit, searchTitleText: searchTitleText)
            switch articles {
            case .success(let articles):
                return .success(articles)
            case .failure(let err):
                return .failure(
                    .init(message: "failed find articles", innerError: err, kind: .internalError))
            }
        }
        
        func insert(title: String, body: String) async -> Result<(), UsecaseError> {
            let result = await self.articleRepo.insert(title: title, body: body)
            
            switch result {
            case .success(_):
                return .success(())
            case .failure(let err):
                return .failure(.init(message: "failed insert article", innerError: err, kind: .internalError))
            }
        }
        
        func update(article: DomainModel.Article) async -> Result<(), UsecaseError> {
            let result = await self.articleRepo.update(article: article)
            switch result {
            case .success(_):
                return .success(())
            case .failure(let err):
                return .failure(.init(message: "failed update article", innerError: err, kind: .internalError))
            }
        }
    }
}

#if DEBUG
    extension Usecase {
        struct ArticleUsecaseMock: ArticleUsecase {
            var getResult: Result<DomainModel.Article, Usecase.UsecaseError>?
            var findResult: Result<[DomainModel.ArticleSummary], Usecase.UsecaseError>?
            var insertResult: Result<(), UsecaseError>?
            var updateResult: Result<(), UsecaseError>?

            func get(id: String) async -> Result<DomainModel.Article, Usecase.UsecaseError> {
                guard let result = self.getResult else {
                    return .failure(
                        .init(message: "unspecfied", innerError: nil, kind: .internalError))
                }
                return result
            }

            func find(offset: Int, limit: Int, searchTitleText: String?) async -> Result<
                [DomainModel.ArticleSummary], Usecase.UsecaseError
            > {
                guard let result = self.findResult else {
                    return .failure(.init(message: "", innerError: nil, kind: .internalError))
                }
                return result
            }
            
            func insert(title: String, body: String) async -> Result<(), UsecaseError> {
                guard let result = self.insertResult else {
                    return .failure(.init(message: "", innerError: nil, kind: .internalError))
                }
                return result
            }
            
            func update(article: DomainModel.Article) async -> Result<(), UsecaseError> {
                guard let result = self.updateResult else {
                    return .failure(.init(message: "", innerError: nil, kind: .internalError))
                }
                return result
            }
        }
    }
#endif

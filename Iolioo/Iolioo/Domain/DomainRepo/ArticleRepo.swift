import Spyable

extension DomainRepo {
    
    @Spyable
    protocol Article {
        func get(id: String) async -> Result<DomainModel.Article, DomainRepo.RepoError>
        func delete(id: String) async -> Result<(), DomainRepo.RepoError>
        func update(article: DomainModel.Article) async -> Result<(), DomainRepo.RepoError>
        func find(offset: Int, limit: Int, searchTitleText: String?) async -> Result<[DomainModel.ArticleSummary], DomainRepo.RepoError>
    }
}

extension DomainRepo {
    protocol Media {
        func find(offset: Int, limit: Int) async -> Result<[DomainModel.MediumSummary], DomainRepo.RepoError>
    }
}

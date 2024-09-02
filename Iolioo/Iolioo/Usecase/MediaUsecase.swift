extension Usecase {
    protocol MediaUsecase {
        func find(offset: Int, limit: Int) async -> Result<[DomainModel.MediumSummary], UsecaseError>
        func get(id: String) async -> Result<DomainModel.Medium, UsecaseError>
    }
}

extension Usecase {
    class MediaUsecaseImpl: MediaUsecase {
        var mediaRepo: DomainRepo.Media
        init(mediaRepo: DomainRepo.Media) {
            self.mediaRepo = mediaRepo
        }
        
        func find(offset: Int, limit: Int) async -> Result<[DomainModel.MediumSummary], Usecase.UsecaseError> {
            let summaries = await self.mediaRepo.find(offset: offset, limit: limit)
            switch summaries {
            case .success(let summaries):
                return .success(summaries)
            case .failure(let err):
                return .failure(.init(message: "failed find media", innerError: err, kind: .internalError))
            }
        }
        
        func get(id: String) async -> Result<DomainModel.Medium, UsecaseError> {
            switch await self.mediaRepo.get(id: id) {
            case .success(let medium):
                return .success(medium)
            case .failure(let err):
                return .failure(.init(message: "failed get medium", innerError: err, kind: .internalError))
            }
        }
    }
}

#if DEBUG
extension Usecase {
    struct MediaUsecaesMock: MediaUsecase {
        var findResult: Result<[DomainModel.MediumSummary], UsecaseError>?
        var getResult: Result<DomainModel.Medium, UsecaseError>?
        func find(offset: Int, limit: Int) async -> Result<[DomainModel.MediumSummary], Usecase.UsecaseError> {
            guard let result = self.findResult else {
                return .failure(.init(message: "", innerError: nil, kind: .internalError))
            }
            return result
        }
        func get(id: String) async -> Result<DomainModel.Medium, UsecaseError> {
            guard let result = self.getResult else {
                return .failure(.init(message: "", innerError: nil, kind: .internalError))
            }
            return result
        }
    }
}
#endif

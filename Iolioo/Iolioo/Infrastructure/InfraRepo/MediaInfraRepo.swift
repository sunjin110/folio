import Foundation
import OpenAPIRuntime
import OpenAPIURLSession

extension InfraRepo {
    
    class Media: DomainRepo.Media {
        let client: APIProtocol
        init(client: APIProtocol) {
            self.client = client
        }
        
        func find(offset: Int, limit: Int) async -> Result<[DomainModel.MediumSummary], DomainRepo.RepoError> {
            
            do {
                let resp = try await self.client.get_sol_media(query: .init(offset: offset, limit: limit))
                let output = try resp.ok.body.json
                
                let summaries = output.media.map {
                    DomainModel.MediumSummary(id: $0.id, fileType: $0.file_type, thumbnailUrl: $0.thumbnail_url, createdAt: $0.created_at, updatedAt: $0.updated_at)
                }
                return .success(summaries)
            } catch {
                return .failure(.init(message: "failed find media. offset: \(offset), limit: \(limit)", innerError: error, kind: .internalError))
            }
        }
    }
}

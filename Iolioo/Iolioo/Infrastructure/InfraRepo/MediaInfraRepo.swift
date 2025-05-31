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
        
        func get(id: String) async -> Result<DomainModel.Medium, DomainRepo.RepoError> {
            do {
                let resp = try await self.client.get_sol_media_sol__lcub_medium_id_rcub_(path: .init(medium_id: id))
                let output = try resp.ok.body.json
                
                return .success(.init(id: output.medium_id, filteType: output.file_type, thumbnailUrl: output.thumbnail_url, downloadUrl: output.download_url))
            } catch {
                return .failure(.init(message: "failed get medium. id: \(id)", innerError: error, kind: .internalError))
            }
        }
    }
}

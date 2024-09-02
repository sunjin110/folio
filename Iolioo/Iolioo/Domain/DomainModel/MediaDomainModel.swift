import Foundation

extension DomainModel {
    struct MediumSummary {
        var id: String
        var fileType: String
        var thumbnailUrl: String?
        var createdAt: Date
        var updatedAt: Date
    }
    
    struct Medium {
        let id: String
        let filteType: String
        let thumbnailUrl: String
        let downloadUrl: String
    }
}

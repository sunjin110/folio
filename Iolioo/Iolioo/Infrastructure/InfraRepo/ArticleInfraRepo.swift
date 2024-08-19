import Foundation
import OpenAPIRuntime
import OpenAPIURLSession

extension InfraRepo {

    func NewArticle(client: APIProtocol) -> DomainRepo.Article {
        return Article.init(client: client)
    }

    class Article: DomainRepo.Article {

        let client: APIProtocol

        init(client: APIProtocol) {
            self.client = client
        }

        public func get(id: String) async -> Result<DomainModel.Article, DomainRepo.RepoError> {
            do {
                let resp = try await self.client.get_sol_articles_sol__lcub_article_id_rcub_(
                    path: .init(article_id: id))
                let output = try resp.ok.body.json
                return .success(
                    DomainModel.Article(
                        id: output.id, title: output.title, body: output.body, writer: "",
                        tags: output.tags.map {
                            DomainModel.ArticleTag.init(id: $0.id, name: $0.name)
                        }, createdAt: output.created_at, updatedAt: output.created_at))
            } catch {
                return .failure(
                    .init(message: "failed request", innerError: error, kind: .internalError))
            }
        }

        public func delete(id: String) async -> Result<(), DomainRepo.RepoError> {
            return .failure(.init(message: "todo", innerError: nil, kind: .internalError))
        }
        
        public func update(article: DomainModel.Article) async -> Result<(), DomainRepo.RepoError> {
            do {
                let resp = try await self.client.put_sol_articles_sol__lcub_article_id_rcub_(
                    path: .init(article_id: article.id),
                    body: .json(
                        .init(
                            title: article.title, body: article.body,
                            tag_ids: article.tags.map { $0.id })))
                let _ = try resp.ok.body.json
                return .success(())
            } catch {
                return .failure(
                    .init(message: "failed request", innerError: error, kind: .internalError))
            }
        }

        public func find(offset: Int, limit: Int, searchTitleText: String?) async -> Result<
            [DomainModel.ArticleSummary], DomainRepo.RepoError
        > {

            do {
                let resp = try await self.client.get_sol_articles(
                    query: .init(offset: offset, limit: limit, search_title_text: searchTitleText))

                let output = try resp.ok.body.json

                let summaries = output.articles.map {
                    DomainModel.ArticleSummary(
                        id: $0.id,
                        title: $0.title,
                        tags: $0.tags.map {
                            DomainModel.ArticleTag(id: $0.id, name: $0.name)
                        },
                        createdAt: ISO8601DateFormatter().date(from: $0.created_at)!,
                        updatedAt: ISO8601DateFormatter().date(from: $0.created_at)!
                    )
                }

                return .success(summaries)
            } catch {
                return .failure(
                    .init(message: "failed find", innerError: error, kind: .internalError))
            }
        }
        
        func insert(title: String, body: String) async -> Result<(), DomainRepo.RepoError> {
            do {
                let resp = try await self.client.post_sol_articles(body: .json(.init(title: title, body: body, tag_ids: [])))
                
                let _ = try resp.ok.body.json
                return .success(())
            } catch {
                return .failure(.init(message: "failed insert article. title: \(title), body: \(body)", innerError: error, kind: .internalError))
            }
        }
    }
}

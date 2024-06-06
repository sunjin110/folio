import { ArticleSummary } from "@/domain/model/article";
import { ArticleRepository, GetArticleSummariesOutput } from "@/domain/repository/article";
import { AuthError, InternalError } from "@/error/error";
import { ArticlesGet200Response, GolioApi, ResponseError } from "@/generate/schema/http";

export function NewArticleRepository(golioApi: GolioApi): ArticleRepository {
    return new article(golioApi);
}

class article implements ArticleRepository {
    private golioApi: GolioApi;

    constructor(golioApi: GolioApi) {
        this.golioApi = golioApi;
    }
    async GenerateBodyByAI(articleID: string, prompt: string): Promise<string> {
        let resp = {
            generatedBody: ""
        };

        try {
            resp = await this.golioApi.articlesArticleIdAiPut({
                articleId: articleID,
                articlesArticleIdAiPutRequest: {
                    message: prompt
                }
            });
        } catch (err) {
            if (err instanceof ResponseError) {
                if (err.response.status === 401 || err.response.status === 403) {
                    throw new AuthError("failed generate body by ai", err);
                }
                throw new InternalError("failed generate body by ai", err);
            }
            throw err;
        }
        return resp.generatedBody;
    }

    async FindSummaries(offset?: number, limit?: number, searchTitleText?: string): Promise<GetArticleSummariesOutput> {
        let resp: ArticlesGet200Response = {
            articles: [],
            total: 0,
        };

        try {
            resp = await this.golioApi.articlesGet({
                offset: offset,
                limit: limit,
                searchTitleText: searchTitleText,
            })
        } catch(err) {
            if (err instanceof ResponseError) {
                if (err.response.status === 401 || err.response.status === 403) {
                    throw new AuthError("failed list articles", err);
                }
                throw new InternalError("failed list articles", err);
            }
            throw err;
        }

        let summaries: ArticleSummary[] = [];
        for (let article of resp.articles) {
            summaries.push({
                id: article.id ?? "",
                title: article.title ?? "",
                created_at: article.createdAt ?? "",
                updated_at: "todo",
            });
        }

        return {
            totalCount: resp.total,
            summaries: summaries,
        };
    }
}
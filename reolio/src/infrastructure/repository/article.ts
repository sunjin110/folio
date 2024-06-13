import { ArticleSummary } from "@/domain/model/article";
import {
  ArticleRepository,
  GetArticleSummariesOutput,
} from "@/domain/repository/article";
import {
  ArticlesAiPost200Response,
  ArticlesGet200Response,
  GolioApi,
} from "@/generate/schema/http";
import { handleError } from "./error";

export function NewArticleRepository(golioApi: GolioApi): ArticleRepository {
  return new article(golioApi);
}

class article implements ArticleRepository {
  private golioApi: GolioApi;

  constructor(golioApi: GolioApi) {
    this.golioApi = golioApi;
  }

  async AsistantBodyByAI(articleID: string, prompt: string): Promise<string> {
    let resp = {
      generatedBody: "",
    };

    try {
      resp = await this.golioApi.articlesArticleIdAiPut({
        articleId: articleID,
        articlesArticleIdAiPutRequest: {
          message: prompt,
        },
      });
    } catch (err) {
      handleError(err, "failed generate body by ai");
    }
    return resp.generatedBody;
  }

  async FindSummaries(
    offset?: number,
    limit?: number,
    searchTitleText?: string,
  ): Promise<GetArticleSummariesOutput> {
    let resp: ArticlesGet200Response = {
      articles: [],
      total: 0,
    };

    try {
      resp = await this.golioApi.articlesGet({
        offset: offset,
        limit: limit,
        searchTitleText: searchTitleText,
      });
    } catch (err) {
      handleError(err, "failed list articles");
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

  async GenerateArticleByAI(prompt: string): Promise<string> {
    let resp: ArticlesAiPost200Response = {
      articleId: "",
    };
    try {
      resp = await this.golioApi.articlesAiPost({
        articlesAiPostRequest: {
          prompt: prompt,
        },
      });
    } catch (err) {
      handleError(err, "failed generate article by ai");
    }
    return resp.articleId;
  }

  async InsertArticle(
    title: string,
    body: string,
    tagIds: string[],
  ): Promise<string> {
    let articleId = "";
    try {
      const resp = await this.golioApi.articlesPost({
        articlesPostRequest: {
          title: title,
          body: body,
          tagIds: tagIds,
        },
      });

      articleId = resp.id;
    } catch (err) {
      handleError(err, "failed insert article");
    }
    return articleId;
  }

  async UpdateArticle(
    articleId: string,
    title: string,
    body: string,
    tagIds: string[],
  ): Promise<string> {
    try {
      await this.golioApi.articlesArticleIdPut({
        articleId: articleId,
        articlesPostRequest: {
          title: title,
          body: body,
          tagIds: tagIds,
        },
      });
    } catch (err) {
      handleError(err, "failed update article");
    }
    return articleId;
  }
}

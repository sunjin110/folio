import { Article, ArticleSummary, ArticleTag } from "@/domain/model/article";
import {
  ArticleRepository,
  GetArticleSummariesOutput,
} from "@/domain/repository/article";
import {
  ArticlesAiPost200Response,
  ArticlesGet200Response,
  GolioApi,
} from "@/generate/schema/http";
import { handleError, wrapError } from "./error";

export function NewArticleRepository(golioApi: GolioApi): ArticleRepository {
  return new article(golioApi);
}

class article implements ArticleRepository {
  private golioApi: GolioApi;

  constructor(golioApi: GolioApi) {
    this.golioApi = golioApi;
  }

  async Get(id: string): Promise<Article> {
    try {
      const resp = await this.golioApi.articlesArticleIdGet({
        articleId: id,
      });
      return {
        id: resp.id,
        title: resp.title,
        body: resp.body,
        tags: resp.tags.map((tag) => {
          return {
            id: tag.id,
            name: tag.name,
          };
        }),
        created_at: resp.createdAt.toISOString(),
      };
    } catch (err) {
      throw wrapError(err, "failed get article");
    }
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
    tagIDs?: string[],
  ): Promise<GetArticleSummariesOutput> {
    let resp: ArticlesGet200Response = {
      articles: [],
      total: 0,
    };

    // get parameterは,で渡さないとダメっぽいのでそれに従う
    const tagIDstr = tagIDs?.join(",");
    let tags: string[] | undefined;
    if (tagIDstr) {
      tags = [tagIDstr];
    }

    try {
      resp = await this.golioApi.articlesGet({
        offset: offset,
        limit: limit,
        searchTitleText: searchTitleText,
        tags: tags,
      });
    } catch (err) {
      handleError(err, "failed list articles");
    }

    let summaries: ArticleSummary[] = [];
    for (let article of resp.articles) {
      const tags: ArticleTag[] = article.tags.map((tag) => ({
        id: tag.id,
        name: tag.name,
      }));

      summaries.push({
        id: article.id ?? "",
        title: article.title ?? "",
        tags: tags,
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

import { Article, ArticleSummary } from "../model/article";

export interface ArticleRepository {
  FindSummaries(
    offset?: number,
    limit?: number,
    searchTitleText?: string,
    tagIDs?: string[],
  ): Promise<GetArticleSummariesOutput>;
  Get(id: string): Promise<Article>;
  AsistantBodyByAI(articleID: string, prompt: string): Promise<string>;
  GenerateArticleByAI(prompt: string): Promise<string>; // 記事ID

  InsertArticle(title: string, body: string, tagIds: string[]): Promise<string>;
  UpdateArticle(
    articleID: string,
    title: string,
    body: string,
    tagIds: string[],
  ): Promise<string>;
}

export interface GetArticleSummariesOutput {
  totalCount: number;
  summaries: ArticleSummary[];
}

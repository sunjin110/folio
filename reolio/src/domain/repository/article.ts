import { ArticleSummary } from "../model/article";

export interface ArticleRepository {
  FindSummaries(
    offset?: number,
    limit?: number,
    searchTitleText?: string,
  ): Promise<GetArticleSummariesOutput>;
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

import { ArticleSummary } from "../model/article";

export interface ArticleRepository {
    FindSummaries(offset?: number, limit?: number, searchTitleText?: string): Promise<GetArticleSummariesOutput>
    GenerateBodyByAI(articleID: string, prompt: string): Promise<string>;
}

export interface GetArticleSummariesOutput {
    totalCount: number;
    summaries: ArticleSummary[];
}

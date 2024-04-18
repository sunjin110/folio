import { ArticleSummary } from "@/domain/model/article";


export interface GetArticleSummariesOutput {
    articles: ArticleSummary[];
    total: number;
};


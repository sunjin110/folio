import { Article, ArticleSummary } from "@/domain/model/article";


export interface GetArticleSummariesOutput {
    articles: ArticleSummary[];
    total: number;
    type: 'success';
};


export interface GetArticleByIdOutput {
    type: 'sucess';
    article: Article;
};
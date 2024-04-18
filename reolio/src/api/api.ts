import { Article, ArticleSummary } from "@/domain/model/article";
import { GetArticleSummariesOutput } from "./response/article";

export async function getArticleById(articleID: string): Promise<Article> {
    const resp = await fetch(process.env.REACT_APP_GOLIO_BASE_URL + `/articles/${articleID}`);
    const article: Article = await resp.json();
    return article;
}

export async function createArticle(title: string, body: string): Promise<void> {
    const resp = await fetch(process.env.REACT_APP_GOLIO_BASE_URL + "/articles", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({'title': title, 'body': body}),
    });
}

export async function getArticles(offset: number, limit: number): Promise<GetArticleSummariesOutput> {
    const resp = await fetch(process.env.REACT_APP_GOLIO_BASE_URL + `/articles?offset=${offset}&limit=${limit}`);
    const output: GetArticleSummariesOutput = await resp.json();
    return output;
}

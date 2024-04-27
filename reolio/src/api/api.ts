import { GetArticleSummariesOutput } from "./response/article";

export interface Article {
    id: string;
    title: string;
    body: string;
    created_at: string;
}

export interface ArticleSummary {
    id: string;
    title: string;
    created_at: string;
    updated_at: string;
}

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

export async function getArticles(): Promise<ArticleSummary[]> {
    const resp = await fetch(process.env.REACT_APP_GOLIO_BASE_URL + "/articles");
    const output: GetArticleSummariesOutput = await resp.json();
    return output.articles;
}

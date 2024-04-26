export interface Article {
    id: string;
    title: string;
    body: string;
    created_at: string;
}

export async function getArticleById(articleID: string): Promise<Article> {
    const resp = await fetch(process.env.REACT_APP_GOLIO_BASE_URL + `/articles/${articleID}`);
    const article: Article = await resp.json();
    return article;
}


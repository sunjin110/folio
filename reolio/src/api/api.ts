import { Article } from "@/domain/model/article";
import {
  CreateArticleOutput,
  GetArticleByIdOutput,
  GetArticleSummariesOutput,
} from "./response/article";

export interface ErrorResponse {
  type: "error";
  message: string;
  status: number;
}

interface EmptyResponse {
  type: "success";
}

export async function getArticleById(
  articleID: string,
): Promise<GetArticleByIdOutput | ErrorResponse> {
  const resp = await fetch(
    process.env.REACT_APP_GOLIO_BASE_URL + `/articles/${articleID}`,
    {
      credentials: "include",
    },
  );
  if (!resp.ok) {
    const msg = await resp.text();
    return {
      type: "error",
      message: msg,
      status: resp.status,
    };
  }
  const article: Article = await resp.json();

  const output: GetArticleByIdOutput = {
    type: "sucess",
    article: article,
  };

  return output;
}

export async function createArticle(
  title: string,
  body: string,
): Promise<CreateArticleOutput | ErrorResponse> {
  const resp = await fetch(process.env.REACT_APP_GOLIO_BASE_URL + "/articles", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ title: title, body: body }),
    credentials: "include",
  });

  if (!resp.ok) {
    const msg = await resp.text();
    return {
      type: "error",
      message: msg,
      status: resp.status,
    };
  }

  const output: CreateArticleOutput = await resp.json();
  return output;
}

export async function updateArticle(
  articleId: string,
  title: string,
  body: string,
): Promise<EmptyResponse | ErrorResponse> {
  const resp = await fetch(
    process.env.REACT_APP_GOLIO_BASE_URL + `/articles/${articleId}`,
    {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ title: title, body: body }),
      credentials: "include",
    },
  );

  if (!resp.ok) {
    const msg = await resp.text();
    return {
      type: "error",
      message: msg,
      status: resp.status,
    };
  }
  return {
    type: "success",
  };
}

export async function getArticles(
  offset: number,
  limit: number,
): Promise<GetArticleSummariesOutput | ErrorResponse> {
  const resp = await fetch(
    process.env.REACT_APP_GOLIO_BASE_URL +
      `/articles?offset=${offset}&limit=${limit}`,
    {
      credentials: "include",
    },
  );
  if (!resp.ok) {
    const msg = await resp.text();
    return {
      type: "error",
      message: msg,
      status: resp.status,
    };
  }
  const output: GetArticleSummariesOutput = await resp.json();
  return output;
}

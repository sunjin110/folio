import { Article, ArticleSummary, ArticleTag } from "@/domain/model/article";
import { ArticleRepository } from "@/domain/repository/article";
import { ArticleTagRepository } from "@/domain/repository/article_tag";

export interface ArticleUsecase {
  FindSummaries(
    offset?: number,
    limit?: number,
    searchTitleText?: string,
  ): Promise<FindArticleSummariesOutput>;
  Get(id: string): Promise<Article>;
  AsistantBodyByAI(articleID: string, prompt: string): Promise<string>; // body
  GenerateArtixleByAI(prompt: string): Promise<string>; // articleID
  InsertArticle(title: string, body: string, tagIds: string[]): Promise<string>;
  UpdateArticle(
    articleID: string,
    title: string,
    body: string,
    tagIds: string[],
  ): Promise<string>;

  FindTags(
    searchText?: string,
    offset?: number,
    limit?: number,
  ): Promise<ArticleTag[]>;

  InsertTag(tagName: string): Promise<string>;
  UpdateTag(tagID: string, tagName: string): Promise<void>;
  DeleteTag(tagID: string): Promise<void>;
}

export interface FindArticleSummariesOutput {
  totalCount: number;
  summaries: ArticleSummary[];
}

export function NewArticleUsecase(
  articleRepo: ArticleRepository,
  articleTagRepo: ArticleTagRepository,
): ArticleUsecase {
  return new article(articleRepo, articleTagRepo);
}

class article implements ArticleUsecase {
  private articleRepo: ArticleRepository;
  private articleTagRepo: ArticleTagRepository;

  constructor(
    articleRepo: ArticleRepository,
    articleTagRepo: ArticleTagRepository,
  ) {
    this.articleRepo = articleRepo;
    this.articleTagRepo = articleTagRepo;
  }

  async Get(id: string): Promise<Article> {
    return await this.articleRepo.Get(id);
  }

  async InsertArticle(
    title: string,
    body: string,
    tagIds: string[],
  ): Promise<string> {
    return await this.articleRepo.InsertArticle(title, body, tagIds);
  }

  async UpdateArticle(
    articleID: string,
    title: string,
    body: string,
    tagIds: string[],
  ): Promise<string> {
    return await this.articleRepo.UpdateArticle(articleID, title, body, tagIds);
  }

  async AsistantBodyByAI(articleID: string, prompt: string): Promise<string> {
    return await this.articleRepo.AsistantBodyByAI(articleID, prompt);
  }

  async FindSummaries(
    offset?: number,
    limit?: number,
    searchTitleText?: string,
  ): Promise<FindArticleSummariesOutput> {
    const output = await this.articleRepo.FindSummaries(
      offset,
      limit,
      searchTitleText,
    );
    return output;
  }

  async GenerateArtixleByAI(prompt: string): Promise<string> {
    return await this.articleRepo.GenerateArticleByAI(prompt);
  }

  async FindTags(
    searchText?: string,
    offset?: number,
    limit?: number,
  ): Promise<ArticleTag[]> {
    return await this.articleTagRepo.Find(searchText, offset, limit);
  }
  async InsertTag(tagName: string): Promise<string> {
    return await this.articleTagRepo.Insert(tagName);
  }
  async UpdateTag(tagID: string, tagName: string): Promise<void> {
    return await this.articleTagRepo.Update(tagID, tagName);
  }
  async DeleteTag(tagID: string): Promise<void> {
    return await this.articleTagRepo.Delete(tagID);
  }
}

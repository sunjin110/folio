import { ArticleSummary } from "@/domain/model/article";
import { ArticleRepository } from "@/domain/repository/article";

export interface ArticleUsecase {
  FindSummaries(
    offset?: number,
    limit?: number,
    searchTitleText?: string,
  ): Promise<FindArticleSummariesOutput>;
  AsistantBodyByAI(articleID: string, prompt: string): Promise<string>; // body
  GenerateArtixleByAI(prompt: string): Promise<string>; // articleID
}

export interface FindArticleSummariesOutput {
  totalCount: number;
  summaries: ArticleSummary[];
}

export function NewArticleUsecase(
  articleRepo: ArticleRepository,
): ArticleUsecase {
  return new article(articleRepo);
}

class article implements ArticleUsecase {
  private articleRepo: ArticleRepository;

  constructor(articleRepo: ArticleRepository) {
    this.articleRepo = articleRepo;
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
}

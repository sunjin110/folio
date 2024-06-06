import { ArticleSummary } from "@/domain/model/article";
import { ArticleRepository } from "@/domain/repository/article";

export interface ArticleUsecase {
    FindSummaries(offset?: number, limit?: number, searchTitleText?: string): Promise<FindArticleSummariesOutput>;
    GenerateBodyByAI(articleID: string, prompt: string): Promise<string>;
}

export interface FindArticleSummariesOutput {
    totalCount: number;
    summaries: ArticleSummary[];
}

export function NewArticleUsecase(articleRepo: ArticleRepository): ArticleUsecase {
    return new article(articleRepo);
}

class article implements ArticleUsecase {

    private articleRepo: ArticleRepository;

    constructor(articleRepo: ArticleRepository) {
        this.articleRepo = articleRepo;
    }

    async GenerateBodyByAI(articleID: string, prompt: string): Promise<string> {
        return await this.articleRepo.GenerateBodyByAI(articleID, prompt);
    }

    async FindSummaries(offset?: number, limit?: number, searchTitleText?: string): Promise<FindArticleSummariesOutput> {
        const output = await this.articleRepo.FindSummaries(offset, limit, searchTitleText);
        return output;
    }
}

import { ArticleTagRepository } from "@/domain/repository/article_tag";
import { GolioApi } from "@/generate/schema/http";
import { handleError } from "./error";
import { ArticleTag } from "@/domain/model/article";

export function NewArticleTagRepository(
  golioApi: GolioApi,
): ArticleTagRepository {
  return new articleTag(golioApi);
}

class articleTag implements ArticleTagRepository {
  private golioApi: GolioApi;

  constructor(golioApi: GolioApi) {
    this.golioApi = golioApi;
  }

  async Find(
    searchText?: string,
    offset?: number,
    limit?: number,
  ): Promise<ArticleTag[]> {
    try {
      const output = await this.golioApi.articleTagsGet({
        searchText: searchText,
        offset: offset,
        limit: limit,
      });
      return output.tags.map((tag) => {
        return {
          id: tag.id,
          name: tag.name,
        };
      });
    } catch (err) {
      handleError(err, "failed find tags");
    }
    return [];
  }

  async Insert(tagName: string): Promise<string> {
    try {
      const output = await this.golioApi.articleTagsPost({
        articleTagsPostRequest: {
          name: tagName,
        },
      });
      return output.id;
    } catch (err) {
      handleError(err, "failed insert tag");
    }
    return "";
  }

  async Delete(tagID: string): Promise<void> {
    try {
      await this.golioApi.articleTagsTagIdDelete({
        tagId: tagID,
      });
    } catch (err) {
      handleError(err, "failed delete tag");
    }
  }

  async Update(tagID: string, tagName: string): Promise<void> {
    try {
      await this.golioApi.articleTagsTagIdPut({
        tagId: tagID,
        articleTagsTagIdPutRequest: {
          name: tagName,
        },
      });
    } catch (err) {
      handleError(err, "failed update tag");
    }
  }
}

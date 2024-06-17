import { ArticleTag } from "../model/article";

export interface ArticleTagRepository {
  Find(
    searchText?: string,
    offset?: number,
    limit?: number,
  ): Promise<ArticleTag[]>;

  Insert(tagName: string): Promise<string>;
  Delete(tagID: string): Promise<void>;
  Update(tagID: string, tagName: string): Promise<void>;
}

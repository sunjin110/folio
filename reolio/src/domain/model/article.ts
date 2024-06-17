export interface Article {
  id: string;
  title: string;
  body: string;
  tags: ArticleTag[];
  created_at: string;
}

export class ArticleSummary {
  id!: string;
  title!: string;
  tags!: ArticleTag[];
  created_at!: string;
  updated_at!: string;
}

export class ArticleTag {
  id!: string;
  name!: string;
}

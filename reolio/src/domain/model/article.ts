export interface Article {
  id: string;
  title: string;
  body: string;
  created_at: string;
}

export class ArticleSummary {
  id!: string;
  title!: string;
  created_at!: string;
  updated_at!: string;
}

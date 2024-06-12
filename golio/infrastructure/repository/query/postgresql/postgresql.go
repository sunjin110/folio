package postgresql

import _ "embed"

var (
	//go:embed upsert_article_summary.sql
	UpsertArticleSummary string

	//go:embed upsert_article_body.sql
	UpsertArticleBody string

	//go:embed upsert_article_tag.sql
	UpsertArticleTag string
)

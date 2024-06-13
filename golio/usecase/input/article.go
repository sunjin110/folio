package input

type ArticleInsert struct {
	Title  string
	Body   string
	TagIDs []string
}

type ArticleUpdate struct {
	ID     string
	Title  string
	Body   string
	TagIDs []string
}

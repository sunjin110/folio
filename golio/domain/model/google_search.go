package model

type GoogleSearchResult struct {
	Title    string
	URL      string
	Overview string
}

type HtmlContent struct {
	Title    string
	URL      string
	Overview string
	BodyText string // tagなどは除外済み
}

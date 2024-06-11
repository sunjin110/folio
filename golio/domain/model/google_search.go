package model

import (
	"fmt"
	"net/url"
)

type GoogleSearchResult struct {
	Title    string
	URL      string
	Overview string
}

func (r *GoogleSearchResult) GetBaseURL() (string, error) {
	u, err := url.ParseRequestURI(r.URL)
	if err != nil {
		return "", fmt.Errorf("failed parse url. url:%s, err: %w", u, err)
	}
	return fmt.Sprintf("%s://%s", u.Scheme, u.Host), nil
}

type HtmlContent struct {
	Title    string
	URL      string
	Overview string
	BodyText string // tagなどは除外済み
}

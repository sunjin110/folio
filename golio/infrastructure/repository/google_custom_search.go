package repository

import (
	"context"
	"fmt"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/gcp/custom_search_api"
)

const (
	// googleSearchEngineID 汎用的なsearchEngineID
	googleSearchEngineID = "8074bbc02f8ef4c26"
)

type googleCustomSearch struct {
	client custom_search_api.Client
}

func NewGoogleCustomSearch(client custom_search_api.Client) repository.GoogleCustomSearch {
	return &googleCustomSearch{
		client: client,
	}
}

func (g *googleCustomSearch) Search(ctx context.Context, searchKeyword string) ([]*model.GoogleSearchResult, error) {
	output, err := g.client.Search(ctx, &custom_search_api.SearchInput{
		SearchText:     searchKeyword,
		SearchEngineID: googleSearchEngineID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed Google Search. err: %w", err)
	}

	searchResults := []*model.GoogleSearchResult{}
	for _, item := range output.Search.Items {
		searchResults = append(searchResults, &model.GoogleSearchResult{
			Title:    item.Title,
			URL:      item.Link,
			Overview: item.Snippet,
		})
	}
	return searchResults, nil
}

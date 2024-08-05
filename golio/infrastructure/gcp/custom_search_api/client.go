package custom_search_api

import (
	"context"
	"fmt"

	"google.golang.org/api/customsearch/v1"
	"google.golang.org/api/option"
)

type Client interface {
	Search(ctx context.Context, input *SearchInput) (*SearchOutput, error)
}

type client struct {
	customsearchService *customsearch.Service
}

func NewClient(ctx context.Context, apiKey string) (Client, error) {
	customsearchService, err := customsearch.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, fmt.Errorf("failed new customsearch Service. err: %w", err)
	}

	return &client{
		customsearchService: customsearchService,
	}, nil
}

type SearchInput struct {
	SearchText     string
	SearchEngineID string  // CX
	FileType       *string // HTML...
	DateRestrict   *string //
}

type SearchOutput struct {
	Search *customsearch.Search
}

func (c *client) Search(ctx context.Context, input *SearchInput) (*SearchOutput, error) {
	call := c.customsearchService.Cse.List().
		Context(ctx).
		Cx(input.SearchEngineID).
		Q(input.SearchText)

	if input.FileType != nil {
		call.FileType(*input.FileType)
	}

	if input.DateRestrict != nil {
		call.DateRestrict(*input.DateRestrict)
	}

	search, err := call.Do()
	if err != nil {
		return nil, fmt.Errorf("failed search. err: %w", err)
	}

	return &SearchOutput{
		Search: search,
	}, nil
}

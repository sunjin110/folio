package repository

import (
	"context"
	"fmt"

	oauth2_v2 "google.golang.org/api/oauth2/v2"
)

type oauth2 struct {
	svc *oauth2_v2.Service
}

func NewOAuth2(ctx context.Context) (*oauth2, error) {

	svc, err := oauth2_v2.NewService(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed oauth2_v2.NewService: %w", err)
	}

	return &oauth2{
		svc: svc,
	}, nil
}

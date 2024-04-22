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

// https://accounts.google.com/o/oauth2/v2/auth?client_id=682633467318-vvlia00uaag3jplkls0uj1md371k54as.apps.googleusercontent.com&redirect_uri=http://localhost:3001/auth/google-oauth/callback&response_type=code&scope=profile email&access_type=offline
func (o *oauth2) Start(ctx context.Context) error {

	// o.svc.Tokeninfo().Context(ctx)

	return nil
}

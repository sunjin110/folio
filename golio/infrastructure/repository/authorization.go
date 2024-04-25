package repository

import (
	"context"
	"net/http"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
)

// https://developers.cloudflare.com/api/
const cloudflareAPIEndpoint = "https://api.cloudflare.com/client/v4"

const (
	readKVPairPath = "/accounts/{:account_id}/storage/kv/namespaces/{:namespace_id}/values/{:key_name}"
)

// https://developers.cloudflare.com/api/operations/workers-kv-namespace-list-namespaces
type authorizationKVStore struct {
	apiToken    string
	namespaceID string
	client      *http.Client
}

func NewAuthorizationKVStore(ctx context.Context, apiToken string, namespaceID string) repository.Authorization {
	return &authorizationKVStore{
		apiToken:    apiToken,
		namespaceID: namespaceID,
		client:      &http.Client{},
	}
}

// CloseSession implements repository.Authorization.
func (a *authorizationKVStore) CloseSession(ctx context.Context, accessToken string) error {
	panic("unimplemented")
}

// Get implements repository.Authorization.
// Subtle: this method shadows the method (*Client).Get of authorizationKVStore.Client.
func (a *authorizationKVStore) Get(ctx context.Context, accessToken string) (*model.UserAuthorization, error) {
	panic("unimplemented")
}

// StartSession implements repository.Authorization.
func (a *authorizationKVStore) StartSession(ctx context.Context, token *model.Token, userAuthorization *model.UserAuthorization) error {
	panic("unimplemented")
}

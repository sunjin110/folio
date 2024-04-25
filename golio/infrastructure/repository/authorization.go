package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"text/template"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto"
)

// https://developers.cloudflare.com/api/
const cloudflareAPIEndpoint = "https://api.cloudflare.com/client/v4"

const (
	readKVPairPath  = "/accounts/{{.AccountID}}/storage/kv/namespaces/{{.NamespaceID}}/values/{{.KeyName}}"
	writeKVPairPath = "/accounts/{{.AccountID}}/storage/kv/namespaces/{{.NamespaceID}}/values/{{.KeyName}}"
)

// https://developers.cloudflare.com/api/operations/workers-kv-namespace-list-namespaces
type authorizationKVStore struct {
	apiToken           string
	accountID          string
	namespaceID        string
	client             *http.Client
	readKVPairPathTmp  *template.Template
	writeKVPairPathTmp *template.Template
}

func NewAuthorizationKVStore(ctx context.Context, apiToken string, accountID string, namespaceID string) (repository.Authorization, error) {
	readKVPairPathTmp, err := template.New("read_kv_pair_path").Parse(readKVPairPath)
	if err != nil {
		return nil, fmt.Errorf("failed make readKVPairPathTmp: %w", err)
	}

	writeKVPairPathTmp, err := template.New("writer_kv_pair_path").Parse(writeKVPairPath)
	if err != nil {
		return nil, fmt.Errorf("failed make writeKVPairPathTmp: %w", err)
	}

	return &authorizationKVStore{
		apiToken:           apiToken,
		accountID:          accountID,
		namespaceID:        namespaceID,
		client:             &http.Client{},
		readKVPairPathTmp:  readKVPairPathTmp,
		writeKVPairPathTmp: writeKVPairPathTmp,
	}, nil
}

// CloseSession implements repository.Authorization.
func (a *authorizationKVStore) CloseSession(ctx context.Context, accessToken string) error {
	panic("unimplemented")
}

// Get implements repository.Authorization.
// Subtle: this method shadows the method (*Client).Get of authorizationKVStore.Client.
func (a *authorizationKVStore) Get(ctx context.Context, accessToken string) (*model.UserAuthorization, error) {

	pathBuf := &bytes.Buffer{}
	a.readKVPairPathTmp.Execute(pathBuf, &dto.AuthorizationGetInput{
		AccountID:   a.accountID,
		NamespaceID: a.namespaceID,
		AccessToken: accessToken,
	})

	url := cloudflareAPIEndpoint + pathBuf.String()

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: url: %s, err: %w", url, err)
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.apiToken))
	req.Header.Set("ContentType", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed a.client. url: %s, err: %w", url, err)
	}
	defer resp.Body.Close()

	respBuf := &bytes.Buffer{}
	if _, err := io.ReadAll(respBuf); err != nil {
		return nil, fmt.Errorf("failed io.ReadAll. err: %w", err)
	}

	userAuthorization := &dto.AuthorizationKVValue{}
	if err := json.Unmarshal(respBuf.Bytes(), userAuthorization); err != nil {
		return nil, fmt.Errorf("failed json.Unmarshal. err: %w", err)
	}
	return userAuthorization.ToModel(), nil
}

// StartSession implements repository.Authorization.
func (a *authorizationKVStore) StartSession(ctx context.Context, token *model.Token, userAuthorization *model.UserAuthorization) error {
	panic("unimplemented")
}

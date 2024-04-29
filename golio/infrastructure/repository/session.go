package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"

	"text/template"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto/kvdto"
)

// https://developers.cloudflare.com/api/
const cloudflareAPIEndpoint = "https://api.cloudflare.com/client/v4"

const (
	curdKVPairPath = "/accounts/{{.AccountID}}/storage/kv/namespaces/{{.NamespaceID}}/values/{{.KeyName}}"
)

// https://developers.cloudflare.com/api/operations/workers-kv-namespace-list-namespaces
type sessionKVStore struct {
	apiToken          string
	accountID         string
	namespaceID       string
	client            *http.Client
	curdKVPairPathTmp *template.Template
}

func NewSessionKVStore(ctx context.Context, apiToken string, accountID string, namespaceID string) (repository.Session, error) {
	curdKVPairPathTmp, err := template.New("curd_kv_pair_path").Parse(curdKVPairPath)
	if err != nil {
		return nil, fmt.Errorf("failed make curdKVPairPathTmp: %w", err)
	}

	return &sessionKVStore{
		apiToken:          apiToken,
		accountID:         accountID,
		namespaceID:       namespaceID,
		client:            &http.Client{},
		curdKVPairPathTmp: curdKVPairPathTmp,
	}, nil
}

func (a *sessionKVStore) Start(ctx context.Context, token *model.Token, userSession *model.UserSession) error {
	userSessionDTO := &dto.SessionKVValue{
		Email:       userSession.Email,
		FirstName:   userSession.FirstName,
		LastName:    userSession.LastName,
		AccessToken: token.AccessToken,
	}

	slog.InfoContext(ctx, "sessionKVStore.Start", "userSessionDTO", userSession, "token", token)
	value, err := json.Marshal(userSessionDTO)
	if err != nil {
		return fmt.Errorf("failed userSessionDTO.MarshalJSON(). dto: %+v, err: %w", userSession, err)
	}

	metadata, err := json.Marshal(kvdto.NewMetadata(&token.ExpireTime))
	if err != nil {
		return fmt.Errorf("failed metadata json marshal: %w", err)
	}

	formData := &bytes.Buffer{}
	multiFormDataWriter := multipart.NewWriter(formData)

	valueFormData, err := multiFormDataWriter.CreateFormField("value")
	if err != nil {
		return fmt.Errorf("failed multiFormDataWriter.CreateFormField(\"value\"): %w", err)
	}
	valueFormData.Write(value)

	metadataFormData, err := multiFormDataWriter.CreateFormField("metadata")
	if err != nil {
		return fmt.Errorf("failed multiFormDataWriter.CreateFormField(\"metadata\"): %w", err)
	}
	metadataFormData.Write(metadata)

	contentType := multiFormDataWriter.FormDataContentType()
	multiFormDataWriter.Close()

	url := a.generateURI(a.curdKVPairPathTmp, &kvdto.PathInput{
		AccountID:   a.accountID,
		NamespaceID: a.namespaceID,
		KeyName:     token.AccessToken,
	})

	req, err := http.NewRequestWithContext(ctx, http.MethodPut, url, formData)
	if err != nil {
		return fmt.Errorf("failed http.NewRequestWithContext. url: %s, err: %w", url, err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.apiToken))
	req.Header.Set("Content-Type", contentType)

	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("a.client.Do. url: %s, err: %w", url, err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed sessionKVStore.Start. url: %s, statusCode: %d, body: %s", url, resp.StatusCode, string(b))
	}
	return nil
}

func (a *sessionKVStore) Close(ctx context.Context, accessToken string) error {
	url := a.generateURI(a.curdKVPairPathTmp, &kvdto.PathInput{
		AccountID:   a.accountID,
		NamespaceID: a.namespaceID,
		KeyName:     accessToken,
	})

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return fmt.Errorf("http.NewRequest: url: %s, err: %w", url, err)
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.apiToken))
	req.Header.Set("Content-Type", "application/json")
	_, err = a.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed a.client. url: %s, err: %w", url, err)
	}
	return nil
}

func (a *sessionKVStore) Get(ctx context.Context, accessToken string) (*model.UserSession, error) {
	url := a.generateURI(a.curdKVPairPathTmp, &kvdto.PathInput{
		AccountID:   a.accountID,
		NamespaceID: a.namespaceID,
		KeyName:     accessToken,
	})
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: url: %s, err: %w", url, err)
	}
	req = req.WithContext(ctx)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.apiToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed a.client. url: %s, err: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed io.ReadAll. err: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed get session. statusCode: %d, err: %s", resp.StatusCode, string(b))
	}

	userSession := &dto.SessionKVValue{}
	if err := json.Unmarshal(b, userSession); err != nil {
		return nil, fmt.Errorf("failed json.Unmarshal. body: %s, err: %w", string(b), err)
	}
	return userSession.ToModel(), nil
}

func (a *sessionKVStore) generateURI(pathTemplate *template.Template, pathInput *kvdto.PathInput) string {
	buf := &bytes.Buffer{}
	pathTemplate.Execute(buf, pathInput)
	return cloudflareAPIEndpoint + buf.String()
}

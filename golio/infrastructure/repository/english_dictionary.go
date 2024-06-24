package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto"
)

const (
	englishDictionaryURLFormat = "https://wordsapiv1.p.rapidapi.com/words/%s"
)

type englishDictionary struct {
	client       *http.Client
	rapidAPIKey  string
	rapidAPIHost string
}

func NewEnglishDictionary(rapidAPIKey string, rapidAPIHost string) repository.EnglishDictionary {
	return &englishDictionary{
		client:       &http.Client{},
		rapidAPIKey:  rapidAPIKey,
		rapidAPIHost: rapidAPIHost,
	}
}

func (e *englishDictionary) GetDetail(ctx context.Context, englishWord string) (*model.WordDetail, error) {

	url := fmt.Sprintf(englishDictionaryURLFormat, englishWord)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed new http request. url: %s, err: %w", url, err)
	}

	req.Header.Add("x-rapidapi-key", e.rapidAPIKey)
	req.Header.Add("x-rapidapi-host", e.rapidAPIHost)

	res, err := e.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed request. err: %w", err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed io.ReadALL. err: %w", err)
	}

	respDTO := &dto.EnglishDictionaryResponse{}
	if err := json.Unmarshal(body, respDTO); err != nil {
		return nil, fmt.Errorf("failed json.Unmarshal. body: %s, err: %w", string(body), err)
	}
	return respDTO.ToWordDetailModel(), nil
}

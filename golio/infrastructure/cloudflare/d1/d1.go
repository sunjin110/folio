package d1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sunjin110/folio/golio/infrastructure/cloudflare/d1/dto"
)

// Client https://developers.cloudflare.com/api/operations/cloudflare-d1-query-database
type Client interface {
	Query(ctx context.Context, input *Input) (*Output, error)
}

type Input struct {
	Params []string
	SQL    string
}

type Output struct {
	Results []interface{}
}

type client struct {
	accountID string
	dbID      string
	apiToken  string
	queryPath string
	client    *http.Client
}

const queryPathFormat = "https://api.cloudflare.com/client/v4/accounts/%s/d1/database/%s/query"

func NewClient(accountID string, dbID string, apiToken string) (Client, error) {

	queryPath := fmt.Sprintf(queryPathFormat, accountID, dbID)

	return &client{
		accountID: accountID,
		dbID:      dbID,
		apiToken:  apiToken,
		queryPath: queryPath,
		client:    &http.Client{},
	}, nil
}

func (c *client) Query(ctx context.Context, input *Input) (*Output, error) {
	queryDTO := dto.QueryInput{
		Params: input.Params,
		SQL:    input.SQL,
	}
	query, err := json.Marshal(queryDTO)
	if err != nil {
		return nil, fmt.Errorf("fialed query json.Marshal. queryDTO: %+v, err: %w", queryDTO, err)
	}

	req, err := http.NewRequest(http.MethodPost, c.queryPath, strings.NewReader(string(query)))
	if err != nil {
		return nil, fmt.Errorf("failed http.NewRequest. url: %s, query: %s, err: %w", c.queryPath, string(query), err)
	}

	req = req.WithContext(ctx)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiToken))
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed c.client.DO. err: %w", err)
	}
	defer resp.Body.Close()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed io.ReadAll. err: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed StatusCode: %d, err: %s", resp.StatusCode, string(respBody))
	}

	respDTO := &dto.Response{}
	if err := json.Unmarshal(respBody, respDTO); err != nil {
		return nil, fmt.Errorf("failed response json.Unmarshal. statusCode: %d, buf: %s, err: %w",
			resp.StatusCode, string(respBody), err)
	}
	return &Output{
		Results: respDTO.GetQueryResult(),
	}, nil
}

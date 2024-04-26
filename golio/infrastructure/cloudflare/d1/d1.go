package d1

import (
	"context"
	"fmt"
	"text/template"
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
	Result   string
	Metadata *Metadata
}

type Metadata struct {
	ChangedDB bool
	LastRowID any
}

type client struct {
	accountID     string
	dbID          string
	apiToken      string
	queryPathTemp *template.Template
}

const queryPath = "https://api.cloudflare.com/client/v4/accounts/{{.AccountID}}/d1/database/{{.DatabaseID}}/query"

func NewClient(accountID string, dbID string, apiToken string) (Client, error) {
	pathTemp, err := template.New("d1_path").Parse(queryPath)
	if err != nil {
		return nil, fmt.Errorf("failed make d1PathTemp: %w", err)
	}
	return &client{
		accountID:     accountID,
		dbID:          dbID,
		apiToken:      apiToken,
		queryPathTemp: pathTemp,
	}, nil
}

func (c *client) Query(ctx context.Context, input *Input) (*Output, error) {

	panic("todo")
}

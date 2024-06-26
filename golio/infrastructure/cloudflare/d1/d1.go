package d1

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/cloudflare/cloudflare-go/v2"
	"github.com/cloudflare/cloudflare-go/v2/d1"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// https://github.com/cloudflare/cloudflare-go/blob/v2.3.0/d1/database_test.go
func NewClient(apiKey string) *cloudflare.Client {
	return cloudflare.NewClient(
		option.WithAPIToken(apiKey),
	)
}

type DB[T any] interface {
	Query(ctx context.Context, sql string, params []string) ([]T, error)
}

type db[T any] struct {
	client    *cloudflare.Client
	accountID string
	dbID      string
}

func NewDB[T any](client *cloudflare.Client, accountID string, dbID string) DB[T] {
	return &db[T]{
		client:    client,
		accountID: accountID,
		dbID:      dbID,
	}
}

func (db *db[T]) Query(ctx context.Context, sql string, params []string) ([]T, error) {
	res, err := db.client.D1.Database.Query(ctx, db.dbID, d1.DatabaseQueryParams{
		AccountID: cloudflare.String(db.accountID),
		Sql:       cloudflare.String(sql),
		Params:    cloudflare.F(params),
	})

	if err != nil {
		return nil, fmt.Errorf("failed cloudflare d1 query. sql: %s, params: %+v, err: %w", sql, params, err)
	}

	results := *res
	result := results[0]

	if !result.Success {
		return nil, fmt.Errorf("failed request. sql: %s, params: %+v, raw: %s", sql, params, result.JSON.RawJSON())
	}

	rowsJSON, err := json.Marshal(result.Results)
	if err != nil {
		return nil, fmt.Errorf("failed json.Marshal. err: %w", err)
	}

	list := make([]T, 0)
	if err := json.Unmarshal(rowsJSON, &list); err != nil {
		return nil, fmt.Errorf("failed json.Unmarshal. err: %w", err)
	}
	return list, nil
}

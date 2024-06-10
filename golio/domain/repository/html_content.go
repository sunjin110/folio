package repository

import (
	"context"
	"io"
)

type HtmlContent interface {
	Get(ctx context.Context, url string) (io.ReadCloser, error)
}

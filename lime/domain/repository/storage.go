package repository

import (
	"context"
	"io"
)

type Strage interface {
	SaveContent(ctx context.Context, reader io.ReadCloser) error
}

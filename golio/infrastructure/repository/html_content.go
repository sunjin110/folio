package repository

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/sunjin110/folio/golio/domain/repository"
)

type htmlContent struct {
	httpClient *http.Client
}

func NewHtmlContent() repository.HtmlContent {
	return &htmlContent{
		httpClient: &http.Client{},
	}
}

func (h *htmlContent) Get(ctx context.Context, url string) (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed new request. err: %w", err)
	}
	resp, err := h.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed httpClient.Do. err: %w", err)
	}

	if resp.StatusCode != 200 {
		resp.Body.Close()
		return nil, fmt.Errorf("failed request. statusCode: %d", resp.StatusCode)
	}
	return resp.Body, nil
}

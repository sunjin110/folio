package model

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type HtmlExtractor struct {
	tokenizer *html.Tokenizer
}

func NewHtmlExtractor(htmlContent string) *HtmlExtractor {
	return &HtmlExtractor{
		tokenizer: html.NewTokenizer(strings.NewReader(htmlContent)),
	}
}

func (he *HtmlExtractor) ExtractText(ctx context.Context) (string, error) {
	sb := &strings.Builder{}
	for {
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("extract text timeout. err: %w", ctx.Err())
		default:
			switch he.tokenizer.Next() {
			case html.ErrorToken:
				err := he.tokenizer.Err()
				if errors.Is(err, io.EOF) {
					return sb.String(), nil
				}
				return "", fmt.Errorf("html.ErrorToken. err: %w", err)
			case html.TextToken:
				sb.WriteString(he.tokenizer.Token().Data)
				sb.WriteString("\n")
			}
		}
	}
}

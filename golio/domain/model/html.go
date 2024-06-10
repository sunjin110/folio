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

	depth := 0
	currentIgnoreTag := ""

	ignoreTagMap := map[string]bool{
		"script": true,
		"style":  true,
		"header": true,
		"iframe": true,
		"head":   true,
	}

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
			case html.StartTagToken:
				if depth > 0 {
					continue
				}

				token := he.tokenizer.Token()
				if ignoreTagMap[token.Data] {
					depth++
					currentIgnoreTag = token.Data
				}

			case html.EndTagToken:
				token := he.tokenizer.Token()

				if depth == 0 || currentIgnoreTag != token.Data {
					continue
				}

				if ignoreTagMap[token.Data] {
					depth--
					currentIgnoreTag = ""
				}

			case html.TextToken:
				if depth > 0 {
					continue
				}

				rawData := he.tokenizer.Token().Data

				text := strings.Join(strings.Fields(rawData), " ")
				if text == "" {
					continue
				}

				sb.WriteString(text)
			}
		}
	}
}

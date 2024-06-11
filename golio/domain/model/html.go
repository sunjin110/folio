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
	baseURL   string // https://xxx.com
}

func NewHtmlExtractor(htmlContent string, baseURL string) *HtmlExtractor {
	return &HtmlExtractor{
		tokenizer: html.NewTokenizer(strings.NewReader(htmlContent)),
		baseURL:   baseURL,
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

					text := strings.Join(strings.Fields(sb.String()), " ")
					return text, nil
				}
				return "", fmt.Errorf("html.ErrorToken. err: %w", err)
			case html.StartTagToken:

				if depth > 0 {
					continue
				}

				token := he.tokenizer.Token()
				if token.Data == "img" || token.Data == "a" {
					src := he.getImageAndLinkSRC(token)
					if src == "" {
						continue
					}

					sb.WriteString(" ")
					sb.WriteString(src)
					sb.WriteString(" ")
					continue
				}

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

			case html.SelfClosingTagToken:
				token := he.tokenizer.Token()
				if token.Data == "img" || token.Data == "a" {
					src := he.getImageAndLinkSRC(token)
					if src == "" {
						continue
					}

					sb.WriteString(" ")
					sb.WriteString(src)
					sb.WriteString(" ")
					continue
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

func (he *HtmlExtractor) getImageAndLinkSRC(token html.Token) string {
	if token.Data == "img" || token.Data == "a" {
		for _, att := range token.Attr {
			if att.Key == "src" {

				if strings.HasPrefix(att.Val, "https://") || strings.HasPrefix(att.Val, "http://") {
					return att.Val
				}

				// htmlをparseしてdomainがない場合は、付け足す
				return he.baseURL + att.Val
			}
		}
	}
	return ""
}

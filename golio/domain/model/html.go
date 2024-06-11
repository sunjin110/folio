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
	inATag    bool
	aURL      string
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
					return sb.String(), nil
				}
				return "", fmt.Errorf("html.ErrorToken. err: %w", err)
			case html.StartTagToken:

				if depth > 0 {
					continue
				}

				token := he.tokenizer.Token()
				if token.Data == "a" {
					he.inATag = true
					he.aURL = he.getURLFromToken(token)
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

				if token.Data == "a" {
					he.inATag = false
					he.aURL = ""
				}

				if ignoreTagMap[token.Data] {
					depth--
					currentIgnoreTag = ""
				}

			case html.SelfClosingTagToken:
				token := he.tokenizer.Token()
				if token.Data == "img" {

					url := he.getURLFromToken(token)

					if strings.Contains(url, "image/png;base64") || strings.Contains(url, "image/jpg;base64") {
						continue
					}

					alt := he.getAltFromToken(token)
					sb.WriteString(fmt.Sprintf("![%s](%s)", alt, url))
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

				if he.inATag && he.aURL != "" {
					text = fmt.Sprintf("[%s](%s)", text, he.aURL)
				}

				sb.WriteString(text)
			}
		}
	}
}

func (he *HtmlExtractor) getURLFromToken(token html.Token) string {
	for _, att := range token.Attr {
		if att.Key == "src" || att.Key == "href" {
			return he.getURIFromSRC(att.Val)
		}
	}
	return ""
}

func (he *HtmlExtractor) getAltFromToken(token html.Token) string {
	for _, att := range token.Attr {
		if att.Key == "alt" {
			return att.Val
		}
	}
	return ""
}

func (he *HtmlExtractor) getURIFromSRC(src string) string {
	if strings.HasPrefix(src, "https://") || strings.HasPrefix(src, "http://") {
		return src
	}

	if strings.HasPrefix(src, "//") {
		return "https:" + src
	}

	return he.baseURL + src
}

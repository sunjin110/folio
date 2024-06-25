package template

import (
	_ "embed"
	"text/template"
)

var (
	//go:embed word_detial.md
	wordDetailMarkdownTemplate string
)

var (
	WordDetailMarkdownTemplate = template.Must(template.New("").Parse(wordDetailMarkdownTemplate))
)

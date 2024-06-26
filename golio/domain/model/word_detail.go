package model

import (
	"bytes"
	"fmt"

	"github.com/sunjin110/folio/golio/domain/model/template"
)

type WordDetail struct {
	Word             string
	Definitions      []*WordDefinition
	Frequency        float64
	PronunciationMap PronunciationMap
}

func (wordDetail *WordDetail) ToMarkdown() (string, error) {
	b := &bytes.Buffer{}
	if err := template.WordDetailMarkdownTemplate.Execute(b, wordDetail); err != nil {
		return "", fmt.Errorf("failed ToMarkdown. wordDetail: %+v, err: %w", wordDetail, err)
	}
	return b.String(), nil
}

type WordDefinition struct {
	Definition   string
	PartOfSpeech string // noun, verb...
	Synonyms     []string
	Antonyms     []string
	Examples     []string
}

type PronunciationMap map[string]string // key, noun, verb or all

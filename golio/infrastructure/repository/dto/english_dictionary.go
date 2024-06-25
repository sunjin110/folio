package dto

import "github.com/sunjin110/folio/golio/domain/model"

type WordDetail struct {
	Word          string                   `json:"word" xml:"word"`
	Results       EnglishDictionaryResults `json:"results" xml:"results"`
	Pronunciation map[string]string        `json:"pronunciation" xml:"-"` // xml unsupport map
	Frequency     float64                  `json:"frequency" xml:"frequency"`
}

func NewWordDetailFromModel(m *model.WordDetail) *WordDetail {
	if m == nil {
		return nil
	}

	results := make([]*EnglishDictionaryResult, 0, len(m.Definitions))
	for _, def := range m.Definitions {
		results = append(results, NewEnglishDictionaryResultFromModel(def))
	}

	return &WordDetail{
		Word:          m.Word,
		Results:       results,
		Pronunciation: m.PronunciationMap,
		Frequency:     m.Frequency,
	}
}

func (resp *WordDetail) ToWordDetailModel() *model.WordDetail {
	if resp == nil {
		return nil
	}

	return &model.WordDetail{
		Word:             resp.Word,
		Definitions:      resp.Results.ToWordDefinitionModels(),
		Frequency:        resp.Frequency,
		PronunciationMap: resp.Pronunciation,
	}
}

type EnglishDictionaryResults []*EnglishDictionaryResult

func (results EnglishDictionaryResults) ToWordDefinitionModels() []*model.WordDefinition {
	models := make([]*model.WordDefinition, 0, len(results))
	for _, result := range results {
		if m := result.ToWordDefinitionModel(); m != nil {
			models = append(models, m)
		}
	}
	return models
}

type EnglishDictionaryResult struct {
	Definition   string   `json:"definition" xml:"definition"`
	PartOfSpeech string   `json:"partOfSpeech" xml:"partOfSpeech"`
	SimilarTo    []string `json:"similarTo" xml:"similarTo"`
	Antonyms     []string `json:"antonyms" xml:"antonyms"`
	Synonyms     []string `json:"synonyms" xml:"synonyms"`
	Examples     []string `json:"examples" xml:"examples"`
}

func NewEnglishDictionaryResultFromModel(m *model.WordDefinition) *EnglishDictionaryResult {
	if m == nil {
		return nil
	}
	return &EnglishDictionaryResult{
		Definition:   m.Definition,
		PartOfSpeech: m.PartOfSpeech,
		SimilarTo:    nil, // TODO
		Antonyms:     m.Antonyms,
		Synonyms:     m.Synonyms,
		Examples:     m.Examples,
	}
}

func (result *EnglishDictionaryResult) ToWordDefinitionModel() *model.WordDefinition {
	if result == nil {
		return nil
	}
	return &model.WordDefinition{
		Definition:   result.Definition,
		PartOfSpeech: result.PartOfSpeech,
		Synonyms:     result.Synonyms,
		Antonyms:     result.Antonyms,
		Examples:     result.Examples,
	}
}

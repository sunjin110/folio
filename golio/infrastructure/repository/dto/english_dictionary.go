package dto

import "github.com/sunjin110/folio/golio/domain/model"

type EnglishDictionaryResponse struct {
	Word          string                   `json:"word"`
	Results       EnglishDictionaryResults `json:"results"`
	Pronunciation map[string]string        `json:"pronunciation"`
	Frequency     float64                  `json:"frequency"`
}

func (resp *EnglishDictionaryResponse) ToWordDetailModel() *model.WordDetail {
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
	Definition   string   `json:"definition"`
	PartOfSpeech string   `json:"partOfSpeech"`
	SimilarTo    []string `json:"similarTo"`
	Antonyms     []string `json:"antonyms"`
	Synonyms     []string `json:"synonyms"`
	Examples     []string `json:"examples"`
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

package conv

import (
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
)

func ToWordDetail(wordDetail *model.WordDetail) openapi.WordDetail {
	definitions := make([]openapi.WordDefinition, 0, len(wordDetail.Definitions))
	for _, d := range wordDetail.Definitions {
		definitions = append(definitions, toWordDefinition(d))
	}

	return openapi.WordDetail{
		Word:        wordDetail.Word,
		Definitions: definitions,
		Frequency:   float32(wordDetail.Frequency),
	}
}

func toWordDefinition(wordDefinition *model.WordDefinition) openapi.WordDefinition {
	return openapi.WordDefinition{
		Definition:   wordDefinition.Definition,
		PartOfSpeech: wordDefinition.PartOfSpeech,
		Synonyms:     wordDefinition.Synonyms,
		Antonyms:     wordDefinition.Antonyms,
		Examples:     wordDefinition.Examples,
	}
}

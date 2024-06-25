package repository

import (
	"context"

	"github.com/sunjin110/folio/golio/domain/model"
)

type Translate interface {
	TranslateText(ctx context.Context, text string, sourceLanguage model.TranslateLanguageCode, targetLanguage model.TranslateLanguageCode) (string, error)
	TranslateWordDetail(ctx context.Context, wordDetail *model.WordDetail, sourceLanguage model.TranslateLanguageCode, targetlanguage model.TranslateLanguageCode) (*model.WordDetailWithTranslation, error)
}

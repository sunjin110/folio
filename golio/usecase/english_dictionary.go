package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
)

type EnglishDictionary interface {
	// GetWordDetailWithTranslation
	// Error: ErrNotFound
	GetWordDetailWithTranslation(ctx context.Context, word string, targetLanguage model.TranslateLanguageCode) (*model.WordDetailWithTranslation, error)
}

type englishDictionary struct {
	translateRepo         repository.Translate
	englishDictionaryRepo repository.EnglishDictionary
}

func NewEnglishDictionary(translateRepo repository.Translate, englishDictionaryRepo repository.EnglishDictionary) EnglishDictionary {
	return &englishDictionary{
		translateRepo:         translateRepo,
		englishDictionaryRepo: englishDictionaryRepo,
	}
}

// GetWordDetailWithTranslation .
func (e *englishDictionary) GetWordDetailWithTranslation(ctx context.Context, word string, targetLanguage model.TranslateLanguageCode) (*model.WordDetailWithTranslation, error) {

	wordDetial, err := e.englishDictionaryRepo.GetDetail(ctx, word)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("failed englishDictionaryRepo.GetDetail. word: %s, err: %w", word, err)
	}

	wordDetailWithTranslation, err := e.translateRepo.TranslateWordDetail(ctx, wordDetial, model.EnglishLanguageCode, targetLanguage)
	if err != nil {
		return nil, fmt.Errorf("failed translateRepo.TranslateWordDetail. err: %w", err)
	}

	// TODO ここの処理は頻発するとコストなのでcacheとかしたい

	return wordDetailWithTranslation, nil
}

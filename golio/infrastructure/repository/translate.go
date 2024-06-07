package repository

import (
	"context"
	"fmt"

	awsSDKTranslate "github.com/aws/aws-sdk-go-v2/service/translate"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	awsTranslate "github.com/sunjin110/folio/golio/infrastructure/aws/translate"
)

type translate struct {
	awsTranslateClient awsTranslate.Client
}

func NewTranslate(awsTranslateClient awsTranslate.Client) repository.Translate {
	return &translate{
		awsTranslateClient: awsTranslateClient,
	}
}

func (t *translate) TranslateText(ctx context.Context, text string, sourceLanguage model.TranslateLanguageCode, targetLanguage model.TranslateLanguageCode) (string, error) {
	output, err := t.awsTranslateClient.TranslateText(ctx, &awsSDKTranslate.TranslateTextInput{
		SourceLanguageCode: sourceLanguage.ToPtrStr(),
		TargetLanguageCode: targetLanguage.ToPtrStr(),
		Text:               &text,
	})
	if err != nil {
		return "", fmt.Errorf("failed awsTranslateClient.TranslateText. err: %w", err)
	}
	return *output.TranslatedText, nil
}

package repository

import (
	"context"
	"encoding/xml"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsSDKTranslate "github.com/aws/aws-sdk-go-v2/service/translate"
	"github.com/aws/aws-sdk-go-v2/service/translate/types"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	awsTranslate "github.com/sunjin110/folio/golio/infrastructure/aws/translate"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto"
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

func (t *translate) TranslateWordDetail(ctx context.Context, wordDetail *model.WordDetail, sourceLanguage model.TranslateLanguageCode, targetlanguage model.TranslateLanguageCode) (*model.WordDetailWithTranslation, error) {
	origin := dto.NewWordDetailFromModel(wordDetail)

	content, err := xml.Marshal(origin)
	if err != nil {
		return nil, fmt.Errorf("failed xml marshal. wordDetail: %+v, err: %w", wordDetail, err)
	}

	output, err := t.awsTranslateClient.TranslateDocument(ctx, &awsSDKTranslate.TranslateDocumentInput{
		Document: &types.Document{
			Content:     content,
			ContentType: aws.String("text/html"),
		},
		SourceLanguageCode: sourceLanguage.ToPtrStr(),
		TargetLanguageCode: targetlanguage.ToPtrStr(),
		Settings:           &types.TranslationSettings{},
		TerminologyNames:   []string{},
	})
	if err != nil {
		return nil, fmt.Errorf("failed TranslateDocument")
	}

	translated := &dto.WordDetail{}
	if err := xml.Unmarshal(output.TranslatedDocument.Content, translated); err != nil {
		return nil, fmt.Errorf("failed xml unmarshal. translated: %s, err: %w", string(output.TranslatedDocument.Content), err)
	}

	return &model.WordDetailWithTranslation{
		Origin:     wordDetail,
		Translated: translated.ToWordDetailModel(),
	}, nil
}

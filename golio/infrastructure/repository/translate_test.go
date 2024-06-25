package repository_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/infrastructure/aws/translate"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
)

// go test -v -count=1 -timeout 30s -run ^Test_translate_TranslateText$ github.com/sunjin110/folio/golio/infrastructure/repository
func Test_translate_TranslateText(t *testing.T) {
	SkipConvey("Test_translate_TranslateText", t, func() {
		ctx := context.Background()

		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile("folio-terraform"), config.WithDefaultRegion("ap-northeast-1"))
		So(err, ShouldBeNil)

		client := translate.NewAWSTranslate(cfg)

		translateRepo := repository.NewTranslate(client)

		translatedText, err := translateRepo.TranslateText(ctx, "hello world", model.EnglishLanguageCode, model.JapaneseLanguageCode)
		So(err, ShouldBeNil)
		So(translatedText, ShouldEqual, "ハローワールド")
	})
}

// go test -v -count=1 -timeout 30s -run ^Test_translate_TranslateWordDetail$ github.com/sunjin110/folio/golio/infrastructure/repository
func Test_translate_TranslateWordDetail(t *testing.T) {
	SkipConvey("Test_translate_TranslateWordDetail", t, func() {
		ctx := context.Background()

		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithSharedConfigProfile("folio-terraform"), config.WithDefaultRegion("ap-northeast-1"))
		So(err, ShouldBeNil)

		client := translate.NewAWSTranslate(cfg)

		translateRepo := repository.NewTranslate(client)

		got, err := translateRepo.TranslateWordDetail(ctx, &model.WordDetail{
			Word: "object",
			Definitions: []*model.WordDefinition{
				{
					Definition:   "the goal intended to be attained (and which is believed to be attainable)",
					PartOfSpeech: "noun",
					Synonyms: []string{
						"aim", "objective", "target",
					},
					Antonyms: []string{},
					Examples: []string{
						"the sole object of her trip was to see her children",
					},
				},
				{
					Definition:   "a tangible and visible entity; an entity that can cast a shadow",
					PartOfSpeech: "noun",
					Examples: []string{
						"it was full of rackets, balls and other objects",
					},
				},
			},
			Frequency:        20.0,
			PronunciationMap: map[string]string{},
		}, model.EnglishLanguageCode, model.JapaneseLanguageCode)
		So(err, ShouldBeNil)
		gotJSON, _ := json.Marshal(got)
		fmt.Println("got is ", string(gotJSON))
	})
}

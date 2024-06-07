package repository_test

import (
	"context"
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

package repository_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
	"github.com/sunjin110/folio/golio/presentation/http/httpconf"
)

// go test -v -count=1 -timeout 30s -run ^Test_englishDictionary_GetDetail_Real$ github.com/sunjin110/folio/golio/infrastructure/repository
func Test_englishDictionary_GetDetail_Real(t *testing.T) {
	SkipConvey("Test_englishDictionary_GetDetail_Real", t, func() {

		conf, err := httpconf.NewConfig()
		So(err, ShouldBeNil)

		repo := repository.NewEnglishDictionary(conf.WordsAPI.RapidAPIKey, conf.WordsAPI.RapidAPIHost)
		detail, err := repo.GetDetail(context.Background(), "object")
		So(err, ShouldBeNil)
		detailJSON, _ := json.Marshal(detail)
		fmt.Println("detail is ", string(detailJSON))

	})
}

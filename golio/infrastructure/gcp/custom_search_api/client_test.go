package custom_search_api_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/golio/infrastructure/gcp/custom_search_api"
)

// go test -v -count=1 -timeout 30s -run ^Test_client_Search_Real$ github.com/sunjin110/folio/golio/infrastructure/gcp/custom_search_api
func Test_client_Search_Real(t *testing.T) {
	SkipConvey("Test_client_Search_Real", t, func() {
		ctx := context.Background()
		client, err := custom_search_api.NewClient(ctx, "")
		So(err, ShouldBeNil)

		output, err := client.Search(ctx, &custom_search_api.SearchInput{
			SearchEngineID: "",
			SearchText:     "今日の晩御飯",
		})
		So(err, ShouldBeNil)

		outputJson, err := json.Marshal(output)
		So(err, ShouldBeNil)
		fmt.Println("outputJson is ", string(outputJson))
	})
}

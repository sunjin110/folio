package d1_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/golio/infrastructure/cloudflare/d1"
)

type ArticleBody struct {
	ID                 string `json:"id"`
	ArticleSummariesID string `json:"article_summaries_id"`
	Body               string `json:"body"`
	CreatedAt          int64  `json:"created_at"`
	UpdatedAt          int64  `json:"updated_at"`
}

// go test -v -count=1 -timeout 30s -run ^Test_db_Query$ github.com/sunjin110/folio/golio/infrastructure/cloudflare/d1
func Test_db_Query(t *testing.T) {
	SkipConvey("Test_db_Query", t, func() {
		client := d1.NewClient("")

		db := d1.NewDB[ArticleBody](client, "", "")

		bodies, err := db.List(context.Background(), "select * from article_bodies;", nil)
		So(err, ShouldBeNil)

		bodiesJSON, _ := json.Marshal(bodies)
		fmt.Println("bodies is ", string(bodiesJSON))

	})
}

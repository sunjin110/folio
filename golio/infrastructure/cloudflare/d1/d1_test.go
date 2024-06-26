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
	Convey("Test_db_Query", t, func() {
		client := d1.NewClient("")

		db := d1.NewDB(client, "36c0ecf2ee9a36a0e0ca24b80b10112d", "d1402245-0e4b-4abc-ac5e-d8b75b61d92b")

		bodies := []ArticleBody{}

		err := db.Query(context.Background(), "select * from article_bodies;", nil, &bodies)
		So(err, ShouldBeNil)

		bodiesJSON, _ := json.Marshal(bodies)
		fmt.Println("bodies is ", string(bodiesJSON))

	})
}

package repository_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/infrastructure/cloudflare/d1"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
	"github.com/sunjin110/folio/golio/presentation/http/httpconf"
)

// go test -v -count=1 -timeout 30s -run ^Test_article_Real$ github.com/sunjin110/folio/golio/infrastructure/repository
func Test_article_Real(t *testing.T) {
	Convey("Test_article_Real", t, func() {
		cfg, err := httpconf.NewConfig()
		So(err, ShouldBeNil)

		ctx := context.Background()
		d1Client, err := d1.NewClient(cfg.D1Database.AccountID, cfg.D1Database.DatabaseID, cfg.D1Database.APIToken)
		So(err, ShouldBeNil)
		articleRepo, err := repository.NewArticle(ctx, d1Client)
		So(err, ShouldBeNil)

		article := &model.Article{
			ID:        "id_1",
			Title:     "title_1",
			Body:      "body_1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err = articleRepo.Insert(ctx, article)
		So(err, ShouldBeNil)

		getArticle, err := articleRepo.Get(ctx, "id_1")
		fmt.Println("getArticle is ", getArticle)
	})
}

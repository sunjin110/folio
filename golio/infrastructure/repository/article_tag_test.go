package repository_test

import (
	"context"
	"fmt"
	"sort"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/golio/domain/model"
	domain_repo "github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto/postgres_dto"
)

func Test_article_tag_Insert_Real(t *testing.T) {
	Convey("Test_article_tag_Insert_Real", t, func() {
		type test struct {
			name   string
			before func(db *sqlx.DB)
			tag    *model.ArticleTag
			after  func(db *sqlx.DB)
		}

		tests := []test{
			{
				name: "Insertできること",
				tag: &model.ArticleTag{
					ID:          "2a4ec3e7-02e1-4c98-8e15-90fb36cc1dc2",
					Name:        "dummy_name",
					CreatedTime: time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC),
					UpdatedTime: time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC),
				},
				after: func(db *sqlx.DB) {
					tag := &postgres_dto.ArticleTag{}
					err := db.Get(tag, "select * from article_tags where id = '2a4ec3e7-02e1-4c98-8e15-90fb36cc1dc2';")
					So(err, ShouldBeNil)

					jsonMatch(tag, &postgres_dto.ArticleTag{
						ID:        "2a4ec3e7-02e1-4c98-8e15-90fb36cc1dc2",
						Name:      "dummy_name",
						CreatedAt: time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC),
					})
				},
			},
		}

		for _, tt := range tests {
			Convey(tt.name, func() {
				db, finish := getTestDB(t)
				defer finish()
				if tt.before != nil {
					tt.before(db)
				}
				articleTagRepo := repository.NewArticleTag(db)
				err := articleTagRepo.Insert(context.Background(), tt.tag)
				So(err, ShouldBeNil)
				if tt.after != nil {
					tt.after(db)
				}
			})
		}
	})
}

// go test -v -count=1 -timeout 30s -run ^Test_article_tag_FindByIDs_Real$ github.com/sunjin110/folio/golio/infrastructure/repository
func Test_article_tag_FindByIDs_Real(t *testing.T) {
	Convey("Test_article_tag_FindByIDs_Real", t, func() {
		type test struct {
			name   string
			before func(db *sqlx.DB, repo domain_repo.ArticleTag)
			ids    []string
			want   []*model.ArticleTag
		}

		tests := []test{
			{
				name: "ちゃんと取得できること",
				before: func(db *sqlx.DB, repo domain_repo.ArticleTag) {

					ids := []string{
						"a5533205-6186-4808-8dfc-81a7a8767b50",
						"047b4017-98a4-4965-82b9-f297eb862ac2",
						"a5598c7f-e715-4345-90ec-34d9c3f0dcf0",
						"59b5f777-31a5-4c27-9700-2fa1184b84cf",
						"914e6519-7832-479e-9ab0-fef2ff44bf47",
						"53b1def1-565d-4e6f-9894-2ec484fa562a",
						"437afb62-2c00-4f6b-8cd1-e72a6fbd29d8",
						"a76dbd68-1ae6-4343-b58b-d6132ad55174",
						"1b9f9a15-6239-4301-ab22-1bbfb4095adc",
						"8fd339bc-6818-4acc-b88d-09f903aef603",
					}

					for i, id := range ids {
						i := i
						err := repo.Insert(context.Background(), &model.ArticleTag{
							ID:          id,
							Name:        fmt.Sprintf("tag_name_%d", i),
							CreatedTime: time.Date(2024, 6, 12, i, 0, 0, 0, time.UTC),
							UpdatedTime: time.Date(2024, 6, 12, i, 0, 0, 0, time.UTC),
						})
						So(err, ShouldBeNil)
					}
				},
				ids: []string{
					"a5533205-6186-4808-8dfc-81a7a8767b50",
					"a5598c7f-e715-4345-90ec-34d9c3f0dcf0",
					"914e6519-7832-479e-9ab0-fef2ff44bf47",
					"437afb62-2c00-4f6b-8cd1-e72a6fbd29d8",
					"1b9f9a15-6239-4301-ab22-1bbfb4095adc",
				},
				want: []*model.ArticleTag{
					{
						ID:          "a5533205-6186-4808-8dfc-81a7a8767b50",
						Name:        "tag_name_0",
						CreatedTime: time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC),
						UpdatedTime: time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:          "a5598c7f-e715-4345-90ec-34d9c3f0dcf0",
						Name:        "tag_name_2",
						CreatedTime: time.Date(2024, 6, 12, 2, 0, 0, 0, time.UTC),
						UpdatedTime: time.Date(2024, 6, 12, 2, 0, 0, 0, time.UTC),
					},
					{
						ID:          "914e6519-7832-479e-9ab0-fef2ff44bf47",
						Name:        "tag_name_4",
						CreatedTime: time.Date(2024, 6, 12, 4, 0, 0, 0, time.UTC),
						UpdatedTime: time.Date(2024, 6, 12, 4, 0, 0, 0, time.UTC),
					},
					{
						ID:          "437afb62-2c00-4f6b-8cd1-e72a6fbd29d8",
						Name:        "tag_name_6",
						CreatedTime: time.Date(2024, 6, 12, 6, 0, 0, 0, time.UTC),
						UpdatedTime: time.Date(2024, 6, 12, 6, 0, 0, 0, time.UTC),
					},
					{
						ID:          "1b9f9a15-6239-4301-ab22-1bbfb4095adc",
						Name:        "tag_name_8",
						CreatedTime: time.Date(2024, 6, 12, 8, 0, 0, 0, time.UTC),
						UpdatedTime: time.Date(2024, 6, 12, 8, 0, 0, 0, time.UTC),
					},
				},
			},
		}

		for _, tt := range tests {
			Convey(tt.name, func() {
				db, finish := getTestDB(t)
				defer finish()
				articleTagRepo := repository.NewArticleTag(db)
				if tt.before != nil {
					tt.before(db, articleTagRepo)
				}
				got, err := articleTagRepo.FindByIDs(context.Background(), tt.ids)
				So(err, ShouldBeNil)

				sort.Slice(got, func(i, j int) bool {
					return got[i].ID < got[j].ID
				})

				sort.Slice(tt.want, func(i, j int) bool {
					return tt.want[i].ID < tt.want[j].ID
				})

				jsonMatch(got, tt.want)
			})
		}
	})
}

func Test_article_tag_Get_Real(t *testing.T) {
	Convey("Test_article_tag_Get_Real", t, func() {
		type test struct {
			name   string
			before func(db *sqlx.DB, repo domain_repo.ArticleTag)
			id     string
			want   *model.ArticleTag
		}

		tests := []test{
			{
				name: "Getできること",
				before: func(db *sqlx.DB, repo domain_repo.ArticleTag) {
					err := repo.Insert(context.Background(), &model.ArticleTag{
						ID:          "662cce87-b30c-422d-8a59-f0adf8212e06",
						Name:        "dummy_name",
						CreatedTime: time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC),
						UpdatedTime: time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC),
					})
					So(err, ShouldBeNil)
				},
				id: "662cce87-b30c-422d-8a59-f0adf8212e06",
				want: &model.ArticleTag{
					ID:          "662cce87-b30c-422d-8a59-f0adf8212e06",
					Name:        "dummy_name",
					CreatedTime: time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC),
					UpdatedTime: time.Date(2024, 6, 12, 0, 0, 0, 0, time.UTC),
				},
			},
		}

		for _, tt := range tests {
			Convey(tt.name, func() {
				db, finish := getTestDB(t)
				defer finish()
				articleTagRepo := repository.NewArticleTag(db)
				if tt.before != nil {
					tt.before(db, articleTagRepo)
				}
				got, err := articleTagRepo.Get(context.Background(), tt.id)
				So(err, ShouldBeNil)

				jsonMatch(got, tt.want)
			})
		}
	})
}

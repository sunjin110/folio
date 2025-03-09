package repository_test

import (
	"context"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/golio/domain/model"
	drepo "github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/domain/repository/mock_repo"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
	"go.uber.org/mock/gomock"
)

// go test -v -count=1 -timeout 30s -run ^Test_articleV2_FindSummary$ github.com/sunjin110/folio/golio/infrastructure/repository
func Test_articleV2_FindSummary(t *testing.T) {
	Convey("Test_articleV2_FindSummary", t, func() {
		type args struct {
			sortType drepo.SortType
			paging   *drepo.Paging
			search   *drepo.ArticleSearch
		}

		type test struct {
			name    string
			args    args
			mocks   func(mock sqlmock.Sqlmock, articleTagRepo *mock_repo.MockArticleTag)
			want    []*model.ArticleSummary
			wantErr error
		}

		tests := []test{
			{
				name: "tag_idsの検索が想定したものであること",
				args: args{
					sortType: drepo.SortTypeAsc,
					paging: &drepo.Paging{
						Offset: 0,
						Limit:  10,
					},
					search: &drepo.ArticleSearch{
						Tags: []string{"tag_1", "tag_2"},
					},
				},
				mocks: func(mock sqlmock.Sqlmock, articleTagRepo *mock_repo.MockArticleTag) {
					rows := sqlmock.NewRows([]string{"id", "title", "tag_ids", "created_at", "updated_at"}).
						AddRow("article_id_1", "title_1", pq.StringArray{"tag_1"}, time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC), time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC)).
						AddRow("article_id_2", "title_2", pq.StringArray{"tag_2"}, time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC), time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC))

					mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM article_summaries WHERE tag_ids @> array[$1,$2] ORDER BY created_at ASC LIMIT $3 OFFSET $4")).
						WithArgs("tag_1", "tag_2", 10, 0).
						WillReturnRows(rows)

					articleTagRepo.EXPECT().FindByIDs(gomock.Any(), gomock.InAnyOrder([]string{"tag_1", "tag_2"})).Return([]*model.ArticleTag{
						{
							ID:          "tag_1",
							Name:        "tag_name_1",
							CreatedTime: time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC),
							UpdatedTime: time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC),
						},
						{
							ID:          "tag_2",
							Name:        "tag_name_2",
							CreatedTime: time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC),
							UpdatedTime: time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC),
						},
					}, nil)
				},
				want: []*model.ArticleSummary{
					{
						ID:     "article_id_1",
						Title:  "title_1",
						Writer: "TODO",
						Tags: []*model.ArticleTag{
							{
								ID:          "tag_1",
								Name:        "tag_name_1",
								CreatedTime: time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC),
								UpdatedTime: time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC),
							},
						},
						CreatedAt: time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:     "article_id_2",
						Title:  "title_2",
						Writer: "TODO",
						Tags: []*model.ArticleTag{
							{
								ID:          "tag_2",
								Name:        "tag_name_2",
								CreatedTime: time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC),
								UpdatedTime: time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC),
							},
						},
						CreatedAt: time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC),
						UpdatedAt: time.Date(2024, 6, 13, 0, 0, 0, 0, time.UTC),
					},
				},
			},
		}

		for _, tt := range tests {
			Convey(tt.name, func() {
				ctx := context.Background()

				ctrl := gomock.NewController(t)
				defer ctrl.Finish()

				db, mock, err := sqlmock.New()
				So(err, ShouldBeNil)
				sqlxDB := sqlx.NewDb(db, "postgres")
				articleTagRepo := mock_repo.NewMockArticleTag(ctrl)

				if tt.mocks != nil {
					tt.mocks(mock, articleTagRepo)
				}

				repo := repository.NewArticleV2(context.Background(), sqlxDB, articleTagRepo)

				got, err := repo.FindSummary(ctx, tt.args.sortType, tt.args.paging, tt.args.search)
				if tt.wantErr != nil {
					So(err, ShouldBeError)
					So(err.Error(), ShouldEqual, tt.wantErr.Error())
					return
				}

				So(err, ShouldBeNil)
				So(got, ShouldResemble, tt.want)
				if tt.mocks != nil {
					So(mock.ExpectationsWereMet(), ShouldBeNil)
				}
			})
		}
	})
}

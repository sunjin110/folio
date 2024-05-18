package repository_test

import (
	"context"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jmoiron/sqlx"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/sunjin110/folio/golio/infrastructure/repository"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto"
)

// go test -v -count=1 -timeout 30s -run ^Test_media_Insert_Real$ github.com/sunjin110/folio/golio/infrastructure/repository
func Test_media_Insert_Real(t *testing.T) {
	Convey("Test_media_Insert_Real", t, func() {
		type args struct {
			id       string
			fileType string
			txTime   time.Time
		}

		type test struct {
			name       string
			before     func(db *sqlx.DB)
			args       args
			after      func(db *sqlx.DB)
			wantPrefix string
		}

		tests := []test{
			{
				name: "mediaの追加に成功すること",
				args: args{
					id:       "522d3916-b6a7-4271-badb-072167a1cdab",
					fileType: "txt",
					txTime:   time.Date(2024, 5, 18, 0, 0, 0, 0, time.UTC),
				},
				wantPrefix: "http://localhost:4566/golio-media/default/522d3916-b6a7-4271-badb-072167a1cdab.txt",
				after: func(db *sqlx.DB) {
					media := []*dto.Medium{}
					err := db.SelectContext(context.Background(), &media, "select * from media;")
					So(err, ShouldBeNil)

					jsonMatch(media, []*dto.Medium{
						{
							ID:        "522d3916-b6a7-4271-badb-072167a1cdab",
							FileType:  "txt",
							Path:      "default/522d3916-b6a7-4271-badb-072167a1cdab.txt",
							CreatedAt: time.Date(2024, 5, 18, 0, 0, 0, 0, time.UTC),
							UpdatedAt: time.Date(2024, 5, 18, 0, 0, 0, 0, time.UTC),
						},
					})
				},
			},
		}

		for _, tt := range tests {
			Convey(tt.name, func() {
				db, finish := getTestDB(t)
				defer finish()

				s3Client := getTestS3Client(t)

				if tt.before != nil {
					tt.before(db)
				}

				mediaRepo := repository.NewMedia(db, testBucketName, s3Client)

				got, err := mediaRepo.Insert(context.Background(), tt.args.txTime, tt.args.id, tt.args.fileType)
				So(err, ShouldBeNil)

				So(got, ShouldNotBeEmpty)
				So(strings.HasPrefix(got, tt.wantPrefix), ShouldBeTrue)

				if tt.after != nil {
					tt.after(db)
				}
			})
		}
	})
}

// go test -v -count=1 -timeout 30s -run ^Test_media_Delete$ github.com/sunjin110/folio/golio/infrastructure/repository
func Test_media_Delete(t *testing.T) {
	Convey("Test_media_Delete", t, func() {
		type test struct {
			name   string
			id     string
			before func(db *sqlx.DB, s3Client *s3.Client)
			after  func(db *sqlx.DB, s3Client *s3.Client)
		}

		tests := []test{
			{
				name: "実行できること",
				id:   "522d3916-b6a7-4271-badb-072167a1cdab",
				before: func(db *sqlx.DB, s3Client *s3.Client) {
					mediaRepo := repository.NewMedia(db, testBucketName, s3Client)
					_, err := mediaRepo.Insert(context.Background(), time.Date(2024, 5, 18, 0, 0, 0, 0, time.Local), "522d3916-b6a7-4271-badb-072167a1cdab", "txt")
					So(err, ShouldBeNil)
				},
				after: func(db *sqlx.DB, s3Client *s3.Client) {
					mediaRepo := repository.NewMedia(db, testBucketName, s3Client)
					c, err := mediaRepo.TotalCount(context.Background())
					So(err, ShouldBeNil)
					So(c, ShouldEqual, int32(0))
				},
			},
		}

		for _, tt := range tests {
			Convey(tt.name, func() {
				db, finish := getTestDB(t)
				defer finish()

				s3Client := getTestS3Client(t)

				if tt.before != nil {
					tt.before(db, s3Client)
				}

				mediaRepo := repository.NewMedia(db, testBucketName, s3Client)

				err := mediaRepo.Delete(context.Background(), tt.id)
				So(err, ShouldBeNil)

				if tt.after != nil {
					tt.after(db, s3Client)
				}
			})
		}
	})
}

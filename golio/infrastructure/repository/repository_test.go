package repository_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jmoiron/sqlx"
	"github.com/sunjin110/folio/golio/infrastructure/postgres"

	. "github.com/smartystreets/goconvey/convey"
)

const testBucketName = "golio-media"

func getTestDB(t *testing.T) (db *sqlx.DB, finish func()) {
	t.Helper()
	datasource := "postgres://golion:golio-password@localhost:5442/golio?sslmode=disable"
	if err := postgres.MigrateDB(datasource); err != nil {
		t.Fatalf("failed test db migrate: %+v", err)
	}
	db, err := postgres.OpenDB(datasource)
	if err != nil {
		t.Fatalf("failed oepn db: %+v", err)
	}
	return db, func() {
		db.Close()
		_ = postgres.MigrateDownDB(datasource)
	}
}

func getTestS3Client(t *testing.T) *s3.Client {
	t.Helper()

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:           "http://localhost:4566", // LocalStackのURLに変更してください
				SigningRegion: "us-east-1",             // 署名リージョンも適宜調整
			}, nil
		})),
	)
	if err != nil {
		t.Fatalf("failed load s3 config: %+v", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	return client
}

func jsonMatch[T any](a, b T) {
	aJSON, err := json.Marshal(a)
	So(err, ShouldBeNil)

	bJSON, err := json.Marshal(b)
	So(err, ShouldBeNil)

	So(string(aJSON), ShouldEqual, string(bJSON))

}

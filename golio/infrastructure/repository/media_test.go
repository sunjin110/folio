package repository_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	. "github.com/smartystreets/goconvey/convey"
)

// go test -v -count=1 -timeout 30s -run ^Test_media_Insert_Real$ github.com/sunjin110/folio/golio/infrastructure/repository
func Test_media_Insert_Real(t *testing.T) {
	SkipConvey("test", t, func() {
		cfg, err := config.LoadDefaultConfig(context.TODO(),
			config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL:           "http://localhost:4566", // LocalStackのURLに変更してください
					SigningRegion: "us-east-1",             // 署名リージョンも適宜調整
				}, nil
			})),
		)
		if err != nil {
			log.Fatalf("unable to load SDK config, %v", err)
		}

		// S3サービスクライアントを作成
		client := s3.NewFromConfig(cfg, func(o *s3.Options) {
			o.UsePathStyle = true
		})

		// プリサイナーを作成
		presigner := s3.NewPresignClient(client)

		// プリサインドURLのリクエストを作成
		presignedReq, err := presigner.PresignGetObject(context.TODO(), &s3.GetObjectInput{
			Bucket: aws.String("golio-media"),
			Key:    aws.String("test.text"),
		}, s3.WithPresignExpires(15*time.Minute)) // 有効期限を15分に設定
		if err != nil {
			log.Fatalf("failed to presign request, %v", err)
		}

		// プリサインドURLを表示
		fmt.Println("Presigned URL:", presignedReq.URL)
	})
}

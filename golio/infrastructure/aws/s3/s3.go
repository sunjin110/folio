package s3

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Client(cfg aws.Config) *s3.Client {
	return s3.NewFromConfig(cfg)
}

func NewPresignClient(cfg aws.Config) *s3.PresignClient {
	s3Client := s3.NewFromConfig(cfg)
	return s3.NewPresignClient(s3Client)
}

//nolint:staticcheck // まだLocalstackが対応してなさそう
func NewLocalStackS3Client() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) { //nolint:staticcheck
			return aws.Endpoint{ //nolint:staticcheck
				URL:           "http://localhost:4566", // LocalStackのURLに変更してください
				SigningRegion: "us-east-1",             // 署名リージョンも適宜調整
			}, nil
		})),
	)
	if err != nil {
		return nil, fmt.Errorf("failed load s3 config: %+v", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	})
	return client, nil
}

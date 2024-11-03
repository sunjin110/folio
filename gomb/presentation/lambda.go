package presentation

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sunjin110/folio/gomb/application"
	"github.com/sunjin110/folio/gomb/config"
	"github.com/sunjin110/folio/gomb/infrastructure/repository"
	"gocv.io/x/gocv"
)

func GetS3EventLambdaEntrypoint(ctx context.Context) (lambdaHandlerFunc func(ctx context.Context, req events.S3Event) error, err error) {

	cfg, err := config.NewEnvConfig()
	if err != nil {
		return nil, fmt.Errorf("failed NewEnvConfig. err: %w", err)
	}

	awsCfg, err := awscfg.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed load aws config. err: %w", err)
	}

	s3Client := s3.NewFromConfig(awsCfg)

	storageRepo := repository.NewStorage(s3Client, cfg.MediaS3BucketName)

	thumbnailApp := application.NewThumbanilApplication(storageRepo)

	return func(ctx context.Context, req events.S3Event) error {
		fmt.Printf("gocv version: %s\n", gocv.Version())
		fmt.Printf("opencv lib version: %s\n", gocv.OpenCVVersion())
		for _, record := range req.Records {
			slog.Info("triggered put event!!", "bucket", record.S3.Bucket, "key", record.S3.Object.Key)
			if err := thumbnailApp.GenerateAndPutThumbnail(ctx, record.S3.Object.Key); err != nil {
				slog.Error("failed generate and put thumbnail", "key", record.S3.Object.Key, "err", err)
				continue
			}
			slog.Info("successed generating and putting thumbnail", "key", record.S3.Object.Key)
		}
		return nil
	}, nil
}

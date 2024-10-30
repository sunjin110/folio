package presentation

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"gocv.io/x/gocv"
)

func GetS3EventLambdaEntrypoint() (lambdaHandlerFunc func(ctx context.Context, req events.S3Event) error, err error) {
	return func(ctx context.Context, req events.S3Event) error {
		slog.Info("========== hello gomb!")
		fmt.Printf("gocv version: %s\n", gocv.Version())
		fmt.Printf("opencv lib version: %s\n", gocv.OpenCVVersion())
		return nil
	}, nil
}

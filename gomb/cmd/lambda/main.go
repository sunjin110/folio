package main

import (
	"context"
	"log/slog"
	"os"

	awslambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/sunjin110/folio/gomb/presentation"
)

func main() {
	entrypoint, err := presentation.GetS3EventLambdaEntrypoint(context.Background())
	if err != nil {
		slog.Error("failed get entrypoint", "err", err)
		os.Exit(1)
	}
	awslambda.Start(entrypoint)
}

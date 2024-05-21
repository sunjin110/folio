package main

import (
	"context"
	"log/slog"
	"os"

	awslambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/sunjin110/folio/golio/presentation/lambda"
)

func main() {
	if err := lambda.Setup(); err != nil {
		slog.Error("failed lambda setup", "err", err)
		os.Exit(1)
	}
	handler, err := lambda.GetHandler(context.Background())
	if err != nil {
		slog.Error("failed get handler", "err", err)
		os.Exit(1)
	}
	awslambda.Start(handler)
}

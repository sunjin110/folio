package main

import (
	"log/slog"

	awslambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/sunjin110/folio/gomb/presentation"
)

func main() {
	entrypoint, err := presentation.GetS3EventLambdaEntrypoint()
	if err != nil {
		slog.Error("failed get entrypoint", "err", err)
		panic(err)
	}
	awslambda.Start(entrypoint)
}

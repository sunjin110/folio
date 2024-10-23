package main

import (
	"log/slog"

	awslambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/sunjin110/folio/lime/presentation"
)

func main() {
	handler, err := presentation.GetLambdaHandler()
	if err != nil {
		slog.Error("failed get handler", "err", err)
		panic(err)
	}
	awslambda.Start(handler)
}

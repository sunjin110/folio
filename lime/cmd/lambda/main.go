package main

import (
	awslambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/sunjin110/folio/lime/presentation"
)

func main() {
	handler, err := presentation.GetHandler()
	if err != nil {
		panic(err)
	}
	awslambda.Start(handler)
}

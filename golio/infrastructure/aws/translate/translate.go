package translate

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/translate"
)

type Client interface {
	TranslateText(ctx context.Context, input *translate.TranslateTextInput) (*translate.TranslateTextOutput, error)
}

type client struct {
	client *translate.Client
}

func NewAWSTranslate(cfg aws.Config) Client {
	return &client{
		client: translate.NewFromConfig(cfg),
	}
}

func (c *client) TranslateText(ctx context.Context, input *translate.TranslateTextInput) (*translate.TranslateTextOutput, error) {
	return c.client.TranslateText(ctx, input)
}

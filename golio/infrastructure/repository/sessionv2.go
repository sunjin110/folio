package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/aws/dynamodb"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto/dynamodto"
)

const (
	indexName = "access_token_index"
)

type sessionV2 struct {
	dynamodbClient dynamodb.Client[dynamodto.UserSessionV2]
	tableName      string
}

func NewSessionV2(dynamodbClient dynamodb.Client[dynamodto.UserSessionV2], tableName string) repository.SessionV2 {
	return &sessionV2{
		dynamodbClient: dynamodbClient,
		tableName:      tableName,
	}
}

func (s *sessionV2) DeleteByAccessToken(ctx context.Context, accessToken string) error {
	userSessionV2, err := s.GetByAccessToken(ctx, accessToken)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			return nil
		}
		return fmt.Errorf("failed get user session. accessToken: %s, err: %w", accessToken, err)
	}

	dto := dynamodto.NewUserSessionV2(userSessionV2)

	key, err := dto.GetKey()
	if err != nil {
		return fmt.Errorf("failed GetKey. err: %w", err)
	}

	if err := s.dynamodbClient.Delete(ctx, s.tableName, key); err != nil {
		return fmt.Errorf("failed delete. err: %w", err)
	}
	return nil
}

func (s *sessionV2) GetByAccessToken(ctx context.Context, accessToken string) (*model.UserSessionV2, error) {
	keyCondition := expression.Key("access_token").Equal(expression.Value(accessToken))

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if err != nil {
		return nil, fmt.Errorf("failed build expression. accessKey: %s, err: %w", accessToken, err)
	}

	limit := int32(1)

	indexName := indexName
	userSessions, _, err := s.dynamodbClient.Query(ctx, s.tableName, expr, &limit, &indexName, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("failed query. accessKey: %s, err: %w", accessToken, err)
	}

	if len(userSessions) == 0 {
		return nil, repository.ErrNotFound
	}
	return userSessions[0].ToModel(), nil
}

func (s *sessionV2) Upsert(ctx context.Context, userSession *model.UserSessionV2) error {
	dto := dynamodto.NewUserSessionV2(userSession)
	if dto == nil {
		return fmt.Errorf("userSession is nil")
	}
	if err := s.dynamodbClient.Add(ctx, s.tableName, *dto); err != nil {
		return fmt.Errorf("failed Upsert userSession. dto: %+v, err: %w", dto, err)
	}
	return nil
}

package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/aws/dynamodb"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto/dynamodto"
)

type sessionV3 struct {
	tableName      string
	dynamodbClient dynamodb.Client[dynamodto.UserSessionV3]
}

func NewSessionV3(dynamodbClient dynamodb.Client[dynamodto.UserSessionV3], tableName string) repository.SessionV3 {
	return &sessionV3{
		tableName:      tableName,
		dynamodbClient: dynamodbClient,
	}
}

func (s *sessionV3) DeleteByEmail(ctx context.Context, email string) error {
	panic("TODO")
}

func (s *sessionV3) Get(ctx context.Context, accessToken string) (*model.UserSessionV3, error) {
	dto := dynamodto.UserSessionV3{
		AccessToken: accessToken,
	}

	key, err := dto.GetKey()
	if err != nil {
		return nil, fmt.Errorf("failed GetKey. err: %w", err)
	}

	userSession, err := s.dynamodbClient.Get(ctx, s.tableName, key)
	if err != nil {
		if errors.Is(err, dynamodb.ErrNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, fmt.Errorf("failed get user_session_v3. err: %w", err)
	}

	return userSession.ToModel(), nil
}

func (s *sessionV3) Upsert(ctx context.Context, userSession *model.UserSessionV3) error {
	if userSession == nil {
		return fmt.Errorf("failed Upsert. userSession is nil")
	}

	dto := dynamodto.NewUserSessionV3(userSession)
	if err := s.dynamodbClient.Add(ctx, s.tableName, *dto); err != nil {
		return fmt.Errorf("failed add userSession. err: ")
	}
	return nil
}

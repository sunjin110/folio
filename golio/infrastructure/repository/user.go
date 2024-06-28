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

type user struct {
	tableName      string
	dynamodbClient dynamodb.Client[dynamodto.User]
}

func NewUser(dynamodbClient dynamodb.Client[dynamodto.User], tableName string) repository.User {
	return &user{
		tableName:      tableName,
		dynamodbClient: dynamodbClient,
	}
}

func (u *user) Get(ctx context.Context, email string) (*model.User, error) {
	dto := dynamodto.User{
		Email: email,
	}
	key, err := dto.GetKey()
	if err != nil {
		return nil, fmt.Errorf("failed GetKey. err: %w", err)
	}

	user, err := u.dynamodbClient.Get(ctx, u.tableName, key)
	if err != nil {
		if errors.Is(err, dynamodb.ErrNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, fmt.Errorf("failed get user. err: %w", err)
	}
	return user.ToModel(), nil
}

func (u *user) Upsert(ctx context.Context, user *model.User) error {
	if user == nil {
		return fmt.Errorf("failed Upsert. user is nil")
	}
	dto := dynamodto.NewUser(user)

	if err := u.dynamodbClient.Add(ctx, u.tableName, *dto); err != nil {
		return fmt.Errorf("failed add user. err: %w", err)
	}
	return nil
}

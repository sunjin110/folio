package dynamodto

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
	"github.com/sunjin110/folio/golio/domain/model"
)

type User struct {
	Email        string `dynamodbav:"email"`
	RefreshToken string `dynamodbav:"refresh_token"`
	FirstName    string `dynamodbav:"first_name"`
	LastName     string `dynamodbav:"last_name"`
	DisplayName  string `dynamodbav:"display_name"`
}

func (dto User) IsDynamoDTO() {}

func (dto User) GetKey() (map[string]types.AttributeValue, error) {
	if dto.Email == "" {
		return nil, errors.New("Email is empty")
	}

	return map[string]types.AttributeValue{
		"email": &types.AttributeValueMemberS{
			Value: dto.Email,
		},
	}, nil
}

func (dto User) ToModel() *model.User {
	return &model.User{
		Email:        dto.Email,
		RefreshToken: dto.RefreshToken,
		FirstName:    dto.FirstName,
		LastName:     dto.LastName,
		DisplayName:  dto.DisplayName,
	}
}

func NewUser(m *model.User) *User {
	if m == nil {
		return nil
	}
	return &User{
		Email:        m.Email,
		RefreshToken: m.RefreshToken,
		FirstName:    m.FirstName,
		LastName:     m.LastName,
		DisplayName:  m.DisplayName,
	}
}

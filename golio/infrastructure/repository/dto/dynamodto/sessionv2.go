package dynamodto

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pkg/errors"
	"github.com/sunjin110/folio/golio/domain/model"
)

type UserSessionV2 struct {
	Email                 string `dynamodbav:"email"`
	AccessToken           string `dynamodbav:"access_token"`
	RefreshToken          string `dynamodbav:"refresh_token"`
	AccessTokenExpireTime int64  `dynamodbav:"access_token_expire_time"` // Unix
	FirstName             string `dynamodbav:"first_name"`
	LastName              string `dynamodbav:"last_name"`
	DisplayName           string `dynamodbav:"display_name"`
}

func (dto UserSessionV2) IsDynamoDTO() {}

func (dto UserSessionV2) GetKey() (map[string]types.AttributeValue, error) {
	if dto.Email == "" {
		return nil, errors.New("Email is empty")
	}

	return map[string]types.AttributeValue{
		"email": &types.AttributeValueMemberS{
			Value: dto.Email,
		},
	}, nil
}

func (dto UserSessionV2) ToModel() *model.UserSessionV2 {
	return &model.UserSessionV2{
		Email:                 dto.Email,
		FirstName:             dto.FirstName,
		LastName:              dto.LastName,
		DisplayName:           dto.DisplayName,
		AccessToken:           dto.AccessToken,
		RefreshToken:          dto.RefreshToken,
		AccessTokenExpireTime: time.Unix(dto.AccessTokenExpireTime, 0).Local(),
	}
}

func NewUserSessionV2(m *model.UserSessionV2) *UserSessionV2 {
	if m == nil {
		return nil
	}
	return &UserSessionV2{
		Email:                 m.Email,
		AccessToken:           m.AccessToken,
		RefreshToken:          m.RefreshToken,
		AccessTokenExpireTime: m.AccessTokenExpireTime.Unix(),
		FirstName:             m.FirstName,
		LastName:              m.LastName,
		DisplayName:           m.DisplayName,
	}
}

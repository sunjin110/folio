package dynamodto

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sunjin110/folio/golio/domain/model"
)

type UserSessionV3 struct {
	AccessToken string `dynamodbav:"access_token"`
	Email       string `dynamodbav:"email"`
	ExpireTime  int64  `dynamodbav:"expire_time"`
}

func (dto UserSessionV3) IsDynamoDTO() {}

func (dto UserSessionV3) GetKey() (map[string]types.AttributeValue, error) {
	if dto.AccessToken == "" {
		return nil, errors.New("AccessToken is empty")
	}

	return map[string]types.AttributeValue{
		"access_token": &types.AttributeValueMemberS{
			Value: dto.AccessToken,
		},
	}, nil
}

func (dto UserSessionV3) ToModel() *model.UserSessionV3 {
	return &model.UserSessionV3{
		AccessToken: dto.AccessToken,
		Email:       dto.Email,
		ExpireTime:  time.Unix(dto.ExpireTime, 0),
	}
}

func NewUserSessionV3(m *model.UserSessionV3) *UserSessionV3 {
	return &UserSessionV3{
		AccessToken: m.AccessToken,
		Email:       m.Email,
		ExpireTime:  m.ExpireTime.Unix(),
	}
}

package dynamodto

import (
	"errors"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sunjin110/folio/golio/domain/model"
)

type UserSessionV3 struct {
	AccessToken           string `dynamodbav:"access_token"`
	Email                 string `dynamodbav:"email"`
	AccessTokenExpireTime int64  `dybamodbav:"access_token_expire_time"`
	ExpireTime            int64  `dynamodbav:"expire_time"` // このデータが自動的に削除される時間
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
		AccessToken:           dto.AccessToken,
		Email:                 dto.Email,
		AccessTokenExpireTime: time.Unix(dto.AccessTokenExpireTime, 0),
	}
}

func NewUserSessionV3(m *model.UserSessionV3) *UserSessionV3 {
	return &UserSessionV3{
		AccessToken:           m.AccessToken,
		Email:                 m.Email,
		AccessTokenExpireTime: m.AccessTokenExpireTime.Unix(),
		ExpireTime:            m.AccessTokenExpireTime.Add(24 * time.Hour).Unix(),
	}
}

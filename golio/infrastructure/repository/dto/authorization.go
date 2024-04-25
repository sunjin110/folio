package dto

import "github.com/sunjin110/folio/golio/domain/model"

// AuthorizationKVValue CloudFlare kvのvalue
type AuthorizationKVValue struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	AccessToken string `json:"access_token"`
}

func (dto *AuthorizationKVValue) ToModel() *model.UserAuthorization {
	if dto == nil {
		return nil
	}
	return &model.UserAuthorization{
		Email:     dto.Email,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}
}

type AuthorizationGetInput struct {
	AccountID   string
	NamespaceID string
	AccessToken string
}

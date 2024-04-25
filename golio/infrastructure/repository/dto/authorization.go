package dto

import (
	"encoding/json"

	"github.com/sunjin110/folio/golio/domain/model"
)

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

func (dto *AuthorizationKVValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(dto)
}

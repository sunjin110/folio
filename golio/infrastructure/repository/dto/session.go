package dto

import (
	"github.com/sunjin110/folio/golio/domain/model"
)

// SessionKVValue CloudFlare kvのvalue
type SessionKVValue struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	AccessToken string `json:"access_token"`
}

func (dto *SessionKVValue) ToModel() *model.UserSession {
	if dto == nil {
		return nil
	}
	return &model.UserSession{
		Email:     dto.Email,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
	}
}

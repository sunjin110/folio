package model

import "time"

type UserSession struct {
	Email       string
	FirstName   string
	LastName    string
	DisplayName string
}

type UserSessionV2 struct {
	Email                 string
	FirstName             string
	LastName              string
	DisplayName           string
	AccessToken           string
	RefreshToken          string
	AccessTokenExpireTime time.Time
}

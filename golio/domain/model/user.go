package model

import "time"

type User struct {
	Email        string
	RefreshToken string
	FirstName    string
	LastName     string
	DisplayName  string
}

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

type UserSessionV3 struct {
	AccessToken string
	Email       string
	ExpireTime  time.Time
}

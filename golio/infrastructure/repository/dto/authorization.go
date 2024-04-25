package dto

// AuthorizationKVValue CloudFlare kvのvalue
type AuthorizationKVValue struct {
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	AccessToken string `json:"access_token"`
}

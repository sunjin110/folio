package dto

type PathInput struct {
	AccountID  string
	DatabaseID string
}

type QueryInput struct {
	Params []string `json:"params"`
	SQL    string   `json:"sql"`
}

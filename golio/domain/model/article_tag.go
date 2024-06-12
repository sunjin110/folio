package model

import "time"

type ArticleTag struct {
	ID          string
	Name        string
	CreatedTime time.Time
	UpdatedTime time.Time
}

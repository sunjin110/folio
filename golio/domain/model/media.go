package model

import "time"

type Medium struct {
	ID           string
	FileType     string
	ThumbnailURL string
	DownloadURL  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type MediumSummary struct {
	ID           string
	FileType     string
	ThumbnailURL string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

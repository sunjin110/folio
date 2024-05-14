package dto

import (
	"time"

	"github.com/sunjin110/folio/golio/domain/model"
)

type Media []*Medium

func (m Media) ToSummariesModel(thumbanilURLMap map[string]string) []*model.MediumSummary {
	summaries := make([]*model.MediumSummary, 0, len(m))
	for _, medium := range m {
		summaries = append(summaries, medium.ToSummaryModel(thumbanilURLMap[medium.ID]))
	}
	return summaries
}

type Medium struct {
	ID        string    `db:"id"`
	Path      string    `db:"path"`
	FileType  string    `db:"file_type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (m *Medium) ToModel(thumbnailURL string, downloadURL string) *model.Medium {
	if m == nil {
		return nil
	}

	return &model.Medium{
		ID:           m.ID,
		FileType:     m.FileType,
		ThumbnailURL: thumbnailURL,
		DownloadURL:  downloadURL,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

func (m *Medium) ToSummaryModel(thumbnailURL string) *model.MediumSummary {
	return &model.MediumSummary{
		ID:           m.ID,
		FileType:     m.FileType,
		ThumbnailURL: thumbnailURL,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
}

package conv

import (
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/generate/schema/http/go/openapi"
)

func ToMediaGet(media []*model.MediumSummary, totalCount int32) openapi.MediaGet200Response {
	return openapi.MediaGet200Response{
		Media:      toMediaGet200ResponseMediaInners(media),
		TotalCount: totalCount,
	}
}

func toMediaGet200ResponseMediaInners(media []*model.MediumSummary) []openapi.MediaGet200ResponseMediaInner {
	inners := make([]openapi.MediaGet200ResponseMediaInner, 0, len(media))
	for _, medium := range media {
		inners = append(inners, toMediaGet200ResponseMediaInner(medium))
	}
	return inners
}

func toMediaGet200ResponseMediaInner(medium *model.MediumSummary) openapi.MediaGet200ResponseMediaInner {
	return openapi.MediaGet200ResponseMediaInner{
		Id:           medium.ID,
		FileType:     medium.FileType,
		ThumbnailUrl: medium.ThumbnailURL,
		CreatedAt:    medium.CreatedAt,
		UpdatedAt:    medium.UpdatedAt,
	}
}

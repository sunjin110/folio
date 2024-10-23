package application

import (
	"context"

	"github.com/sunjin110/folio/lime/domain/model"
)

type LineUsecase interface {
	SaveContents(ctx context.Context, events model.LineEvents) error
}

type lineUsecase struct {
}

func NewLineUsecase() LineUsecase {
	return &lineUsecase{}
}

func (l *lineUsecase) SaveContents(ctx context.Context, events model.LineEvents) error {
	// TODO コンテンツを取得

	// TODO s3に保存

	// TODO say success

	panic("unimplemented")
}

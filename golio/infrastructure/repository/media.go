package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/huandu/go-sqlbuilder"
	"github.com/jmoiron/sqlx"
	"github.com/sunjin110/folio/golio/domain/model"
	"github.com/sunjin110/folio/golio/domain/repository"
	"github.com/sunjin110/folio/golio/infrastructure/repository/dto"
	"github.com/sunjin110/folio/golio/utils/smap"
	"golang.org/x/sync/errgroup"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const mediaPresignedExpireDuration = 10 * time.Minute

type media struct {
	db              *sqlx.DB
	s3BucketName    string
	s3Client        *s3.Client
	presignedClient *s3.PresignClient
}

func NewMedia(db *sqlx.DB, s3BucketName string, s3Client *s3.Client) repository.Media {
	return &media{
		db:              db,
		s3BucketName:    s3BucketName,
		s3Client:        s3Client,
		presignedClient: s3.NewPresignClient(s3Client),
	}
}

func (m *media) Delete(ctx context.Context, id string) error {
	rows, err := m.db.QueryxContext(ctx, `delete from media where id = $1 returning *;`, id)
	if err != nil {
		return fmt.Errorf("failed delete. id: %s, err: %w", id, err)
	}

	defer rows.Close()

	if !rows.Next() {
		return nil
	}

	medium := &dto.Medium{}

	if err := rows.StructScan(&medium); err != nil {
		return fmt.Errorf("failed rows.StructScan: %w", err)
	}

	if _, err := m.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &m.s3BucketName,
		Key:    &medium.Path,
	}); err != nil {
		return fmt.Errorf("failed delete from s3. path: %s, err: %w", medium.Path, err)
	}

	return nil
}

func (m *media) FindSummary(ctx context.Context, paging *repository.Paging) ([]*model.MediumSummary, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("*").From("media").Limit(paging.Limit).Offset(paging.Offset).OrderBy("created_at")
	sb.Desc()

	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	media := dto.Media{}
	if err := m.db.SelectContext(ctx, &media, sql, args...); err != nil {
		return nil, fmt.Errorf("failed findSummary. sql: %s, args: %+v, err: %w", sql, args, err)
	}

	presignedURLMap := smap.NewMap[string, string]()
	eg, ctx := errgroup.WithContext(ctx)
	for _, medium := range media {
		medium := medium
		eg.Go(func() error {
			path := m.generatePath(medium.ID, medium.FileType)
			presignedURL, err := m.getDownloadPresignedURL(ctx, path)
			if err != nil {
				return fmt.Errorf("failed get download presigned url. path: %s, err: %w", path, err)
			}

			presignedURLMap.Put(medium.ID, presignedURL)
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return nil, fmt.Errorf("failed get presigned url. err: %w", err)
	}
	return media.ToSummariesModel(presignedURLMap.GetRawMap()), nil
}

func (m *media) Get(ctx context.Context, id string) (*model.Medium, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("*").From("media").Limit(1).Where(sb.EQ("id", id))
	sql, args := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	mediumDto := &dto.Medium{}
	if err := m.db.GetContext(ctx, mediumDto, sql, args...); err != nil {
		return nil, fmt.Errorf("failed get medium. id: %s, err: %w", id, err)
	}

	downloadPresignedURL, err := m.getDownloadPresignedURL(ctx, mediumDto.Path)
	if err != nil {
		return nil, fmt.Errorf("failed get download presigned url: %w", err)
	}

	// TODO thumbnailは別で発行できるようにする
	return mediumDto.ToModel(downloadPresignedURL, downloadPresignedURL), nil
}

func (m *media) Insert(ctx context.Context, txTime time.Time, id string, fileType string) (uploadPresignedURL string, err error) {
	path := m.generatePath(id, fileType)
	uploadPresignedURL, err = m.getUploadPresignedURL(ctx, path)
	if err != nil {
		return "", fmt.Errorf("failed make upload presigned url: %w", err)
	}
	sb := sqlbuilder.NewInsertBuilder()
	sql, args := sb.InsertInto("media").
		Cols("id", "path", "file_type", "created_at", "updated_at").
		Values(id, path, fileType, txTime, txTime).BuildWithFlavor(sqlbuilder.PostgreSQL)

	if _, err := m.db.ExecContext(ctx, sql, args...); err != nil {
		return "", fmt.Errorf("failed insert to db. sql: %s, err: %w", sql, err)
	}
	return uploadPresignedURL, nil
}

func (m *media) generatePath(id string, fileType string) string {
	// TODO user_idができたらそれごとに分けるようにする
	return fmt.Sprintf("default/%s.%s", id, fileType)
}

func (m *media) getUploadPresignedURL(ctx context.Context, path string) (string, error) {
	req, err := m.presignedClient.PresignPutObject(ctx, &s3.PutObjectInput{
		Bucket: &m.s3BucketName,
		Key:    &path,
	}, func(po *s3.PresignOptions) {
		po.Expires = mediaPresignedExpireDuration
	})
	if err != nil {
		return "", fmt.Errorf("failed make presigned url: %w", err)
	}
	return req.URL, nil
}

func (m *media) getDownloadPresignedURL(ctx context.Context, path string) (string, error) {
	req, err := m.presignedClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: &m.s3BucketName,
		Key:    &path,
	}, func(po *s3.PresignOptions) {
		po.Expires = mediaPresignedExpireDuration
	})
	if err != nil {
		return "", fmt.Errorf("failed make presigned url: %w", err)
	}
	return req.URL, nil
}

func (m *media) TotalCount(ctx context.Context) (int32, error) {
	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("count(*)").From("media")
	sql, _ := sb.BuildWithFlavor(sqlbuilder.PostgreSQL)

	rows, err := m.db.DB.QueryContext(ctx, sql)
	if err != nil {
		return -1, fmt.Errorf("failed queryContext: %w", err)
	}
	defer rows.Close()
	if !rows.Next() {
		return -1, fmt.Errorf("rows not found")
	}

	var totalCount int32
	if err := rows.Scan(&totalCount); err != nil {
		return -1, fmt.Errorf("failed scan: %w", err)
	}
	return totalCount, nil
}

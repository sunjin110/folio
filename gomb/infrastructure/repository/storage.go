package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sunjin110/folio/gomb/domain/model"
	"github.com/sunjin110/folio/gomb/domain/repository"
)

type storage struct {
	s3Client *s3.Client
	bucket   string
}

func NewStorage(s3Client *s3.Client, bucket string) repository.Storage {
	return &storage{
		s3Client: s3Client,
		bucket:   bucket,
	}
}

func (s *storage) GetContent(ctx context.Context, path string) (*model.Content, error) {
	output, err := s.s3Client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return nil, fmt.Errorf("failed GetObject. %w", err)
	}

	var contentType string
	if output.ContentType != nil {
		contentType = *output.ContentType
	}

	f := strings.Split(path, "/")
	fileName := f[len(f)-1]

	return model.NewContent(
		output.Body,
		contentType,
		*output.ContentLength,
		fileName,
	), nil
}

func (s *storage) SaveContent(ctx context.Context, dir string, content *model.Content) error {
	defer content.ReadCloser().Close()

	if !strings.HasSuffix(dir, "/") {
		dir = dir + "/"
	}
	key := dir + content.FileName()

	if _, err := s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(key),
		Body:          content.ReadCloser(),
		ContentLength: aws.Int64(content.ContentLength()),
		ContentType:   aws.String(content.ContentType()),
	}); err != nil {
		return fmt.Errorf("fialed s3Client.PutObject. err: %w", err)
	}
	return nil
}

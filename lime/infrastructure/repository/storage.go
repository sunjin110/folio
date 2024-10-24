package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sunjin110/folio/lime/domain/model"
	"github.com/sunjin110/folio/lime/domain/repository"
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

func (s *storage) SaveContent(ctx context.Context, content *model.Content) (err error) {
	defer content.ReadCloser().Close()
	key, err := s.generateObjectKey(content)
	if err != nil {
		return fmt.Errorf("failed s.generateObjectKey. err: %w", err)
	}

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

func (s *storage) generateObjectKey(content *model.Content) (string, error) {
	fileName, err := content.FileName()
	if err != nil {
		return "", fmt.Errorf("failed content.FileName. err: %w", err)
	}

	return fmt.Sprintf("lime/%s", fileName), nil
}

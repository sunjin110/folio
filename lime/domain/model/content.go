package model

import (
	"fmt"
	"io"
	"log/slog"
	"mime"

	"github.com/gofrs/uuid/v5"
)

type Content struct {
	readCloser    io.ReadCloser
	contentType   string
	contentLength int64
	fileName      *string
}

func (c *Content) ReadCloser() io.ReadCloser {
	return c.readCloser
}

func (c *Content) ContentType() string {
	return c.contentType
}

func (c *Content) FileName() (string, error) {
	if c.fileName != nil {
		return *c.fileName, nil
	}

	// TODO content生成時にする
	u, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("failed uuid.NewV4. err: %w", err)
	}

	s, err := mime.ExtensionsByType(c.contentType)
	if err != nil {
		return "", fmt.Errorf("failed mime.ExtensionsByType. contentType: %s, err: %w", c.contentType, err)
	}

	var ext string
	if len(s) > 0 {
		ext = s[0]
	} else {
		slog.Warn("no extension")
	}

	fileName := fmt.Sprintf("%s%s", u.String(), ext)
	c.fileName = &fileName
	return fileName, nil
}

func (c *Content) ContentLength() int64 {
	return c.contentLength
}

func NewContent(readCloser io.ReadCloser, contentType string, contentLength int64, fileName *string) *Content {

	return &Content{
		readCloser:    readCloser,
		contentType:   contentType,
		contentLength: contentLength,
		fileName:      fileName,
	}
}

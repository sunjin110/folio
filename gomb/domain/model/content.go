package model

import (
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

type Content struct {
	readCloser    io.ReadCloser
	contentType   string
	contentLength int64
	fileName      string
}

func (c *Content) ReadCloser() io.ReadCloser {
	return c.readCloser
}

func (c *Content) ContentType() string {
	return c.contentType
}

func (c *Content) ContentLength() int64 {
	return c.contentLength
}

func (c *Content) FileName() string {
	return c.fileName
}

func (c *Content) getMime() *mimetype.MIME {
	return mimetype.Lookup(c.contentType)
}

func (c *Content) IsImage() bool {
	mime := c.getMime()
	return strings.HasPrefix(mime.String(), "image")
}

func (c *Content) IsVideo() bool {
	mime := c.getMime()
	return strings.HasPrefix(mime.String(), "video")
}

func (c *Content) Thumbnail() (*Image, error) {
	if c.IsImage() {
		i, err := NewImage(c.readCloser)
		if err != nil {
			return nil, fmt.Errorf("failed NewImage. err: %w", err)
		}
		thumbnail := i.Thumbnail()
		return &thumbnail, nil
	} else if c.IsVideo() {
		v, err := NewVideo(c.readCloser)
		if err != nil {
			return nil, fmt.Errorf("failed NewVideo. err: %w", err)
		}
		thumbnail, err := v.Thumbnail()
		if err != nil {
			return nil, fmt.Errorf("failed generate thumbnail from video. err: %w", err)
		}
		return &thumbnail, nil
	}
	return nil, errors.New("this content is not video or image")
}

func NewContent(readCloser io.ReadCloser, contentType string, contentLength int64, fileName string) *Content {
	return &Content{
		readCloser:    readCloser,
		contentType:   contentType,
		contentLength: contentLength,
		fileName:      fileName,
	}
}

package model

import "io"

type Content struct {
	readCloser    io.ReadCloser
	mediatype     string
	contentLength int64
}

func NewContent(readCloser io.ReadCloser, mediatype string, contentLength int64) *Content {
	return &Content{
		readCloser:    readCloser,
		mediatype:     mediatype,
		contentLength: contentLength,
	}
}

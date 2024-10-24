package model

type LineEvents []LineEvent

type LineEvent interface {
	IsLineEvent()
}

type LineMessageEvent struct {
	LineMessageContent *LineMessageContent
}

func (m *LineMessageEvent) IsLineEvent() {}

type LineMessageContent struct {
	ID                  string
	LineContentProvider LineContentProvider
}

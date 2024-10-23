package model

type LineContentProviderType int

const (
	LineContentProviderTypeLine     = 1
	LineContentProviderTypeExternal = 2
)

// LineContentProvider Lineのコンテンツをどこから取得できるか?
type LineContentProvider interface {
	GetType() LineContentProviderType
}

type LineContentProviderExternal struct {
	OriginalContentURL string
	PreviewImageURL    string
}

func (*LineContentProviderExternal) GetType() LineContentProviderType {
	return LineContentProviderTypeExternal
}

type LineContentProviderLine struct {
	MessageID string
}

func (*LineContentProviderLine) GetType() LineContentProviderType {
	return LineContentProviderTypeLine
}

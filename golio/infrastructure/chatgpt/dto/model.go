package dto

type Message interface {
	IsMessage()
}

type SystemMessage struct {
	Role    string `json:"role"` // system
	Content string `json:"content"`
}

func (s *SystemMessage) IsMessage() {}

type UserMessage struct {
	Role    string `json:"role"` // user
	Content string `json:"content"`
}

func (u *UserMessage) IsMessage() {}

type GeneratedMessage struct {
	Role    string  `json:"role"`
	Content *string `json:"content"`
}

type Choise struct {
	FinishReason string            `json:"finish_reason"` // stop | length |content_filter | tool_calls
	Index        int               `json:"index"`
	Message      *GeneratedMessage `json:"message"`
}

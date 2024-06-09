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

type ToolMessage struct {
	Role       string `json:"role"`    // tool
	Content    string `json:"content"` // json
	ToolCallID string `json:"tool_call_id"`
}

func (f *ToolMessage) IsMessage() {}

type AssistantMessage struct {
	Role      string                     `json:"role"` // assistant
	Content   string                     `json:"content"`
	ToolCalls []AssistantMessageToolCall `json:"tool_calls"`
}

func (a *AssistantMessage) IsMessage() {}

type AssistantMessageToolCall struct {
	ID       string                           `json:"id"`
	Type     string                           `json:"type"` // function
	Function AssistantMessageToolCallFunction `json:"function"`
}

type AssistantMessageToolCallFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type GeneratedMessage struct {
	Role      string            `json:"role"`
	Content   *string           `json:"content"`
	ToolCalls []*ChoiceToolCall `json:"tool_calls"`
}

type Choise struct {
	FinishReason string            `json:"finish_reason"` // stop | length |content_filter | tool_calls
	Index        int               `json:"index"`
	Message      *GeneratedMessage `json:"message"`
}

type ChoiceToolCall struct {
	ID       string                  `json:"id"`
	Type     string                  `json:"type"`
	Function *ChoiceToolCallFunction `json:"function"`
}

type ChoiceToolCallFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

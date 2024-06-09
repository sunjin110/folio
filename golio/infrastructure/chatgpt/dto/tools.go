package dto

type Tools []Tool

type Tool interface {
	IsTool()
}

type ToolFunction struct {
	Type     string    `json:"type"` // function
	Function *Function `json:"function"`
}

func (tf *ToolFunction) IsTool() {}

type Function struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Parameters  FuncitonParameters `json:"parameters"`
}

type FuncitonParameters struct {
	Type       string                              `json:"type"` // object
	Properties map[string]*FunctionPropertiesValue `json:"properties"`
	Required   []string                            `json:"required"`
}

type FunctionPropertiesValue struct {
	Type        string `json:"type"` // stringとか
	Description string `json:"description"`
}

package dto

type Response struct {
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
	Results  []*Result     `json:"result"`
	Success  bool          `json:"success"`
}

func (resp *Response) GetQueryResult() []interface{} {
	if resp == nil {
		return nil
	}

	for _, result := range resp.Results {
		if len(result.Results) > 0 {
			return result.Results
		}
	}
	return nil
}

func (resp *Response) GetMeta() *Meta {
	if resp == nil {
		return nil
	}
	for _, result := range resp.Results {
		if result.Meta != nil {
			return result.Meta
		}
	}
	return nil
}

type Result struct {
	Meta    *Meta         `json:"meta"`
	Results []interface{} `json:"results"`
	Success bool          `json:"success"`
}

type Meta struct {
	ChangedDB bool    `json:"changed_db"`
	Changes   int     `json:"changes"`
	Duration  float64 `json:"duration"`
}

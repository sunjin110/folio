package dto

type Response struct {
	Errors   []interface{} `json:"errors"`
	Messages []interface{} `json:"messages"`
	Results  []*Result     `json:"result"`
	Success  bool          `json:"success"`
}

type Result struct {
	Meta    *Meta         `json:"meta"`
	Results []interface{} `json:"results"`
	Success bool          `json:"success"`
}

type Meta struct {
	ChangedDB bool `json:"changed_db"`
	Changes   int  `json:"changes"`
	Duration  int  `json:"duration"`
}

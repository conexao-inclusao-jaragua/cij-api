package model

type Response struct {
	Message string      `json:"message"`
	Code    string      `json:"code,omitempty"`
	Fields  []string    `json:"fields,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type LoginResponse struct {
	Token    string      `json:"token,omitempty"`
	UserInfo interface{} `json:"user_info,omitempty"`
	Message  string      `json:"message,omitempty"`
}

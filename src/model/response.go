package model

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

type LoginResponse struct {
	Token    string `json:"token"`
	UserInfo User   `json:"user_info"`
}

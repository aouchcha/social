package entity

type Response struct {
	Msg     string  `json:"msg"`
	Code    int     `json:"code"`
	Session Session `json:"token,omitempty"`
	UserID  uint    `json:"userId,omitempty"`
}

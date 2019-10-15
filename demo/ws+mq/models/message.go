package models

type Message struct {
	ID     uint                   `json:"id"`
	UserID uint                   `json:"user_id"`
	Msg    map[string]interface{} `json:"msg"`
}

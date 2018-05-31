package model

type MsgEvent struct {
	Kind    int         `json:"kind"`
	Content interface{} `json:"content"`
}

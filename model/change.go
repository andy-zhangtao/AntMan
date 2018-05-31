package model

type Change struct {
	Name    string   `json:"name"`
	Address []string `json:"address"`
}

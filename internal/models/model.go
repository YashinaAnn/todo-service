package models

type Todo struct {
	ID        int64  `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

package models

type Item struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Price  int    `json:"price"`
}

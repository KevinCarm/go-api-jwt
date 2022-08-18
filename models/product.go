package models

type Product struct {
	Id     int     `json:"id"`
	UserId int     `json:"userId"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
}

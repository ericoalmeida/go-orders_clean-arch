package domain

import "time"

type Order struct {
	ID           string    `json:"id"`
	Item         string    `json:"item"`
	Customer     string    `json:"customer"`
	PurchaseDate time.Time `json:"purchaseDate"`
	Price        int64     `json:"price"`
}

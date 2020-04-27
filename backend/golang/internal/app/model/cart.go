package model

import "time"

type Cart struct {
	ID        int `json:"id"`
	Customer  Customer
	CreatedAt time.Time `json:"created_at"`
}

type CartProduct struct {
	Customer Customer
	Product  Product
}

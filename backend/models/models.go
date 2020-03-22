package models

import "time"

// Delivery ...
type Delivery struct {
	ID                  uint64    `json:"id,omitempty"`
	StartAt             time.Time `json:"startAt,omitempty"`
	FinalAt             time.Time `json:"finalAt,omitempty"`
	TotalWeightProducts int       `json:"totalWeightProducts,omitempty"`
	Man                 Man       `json:"man,omitempty"`
	Products            []Product `json:"products,omitempty"`
}

// Product ...
type Product struct {
	ID      uint64 `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Weight  int    `json:"weight,omitempty"`
	BarCode int    `json:"barCode,omitempty"`
}

// Man ...
type Man struct {
	ID          uint64      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	LastName    string      `json:"lastName,omitempty"`
	PhoneNumber string      `json:"phoneNumber,omitempty"`
	Transport   []Transport `json:"transport,omitempty"`
}

// Transport ...
type Transport struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

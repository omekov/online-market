package models

import "time"

type Delivery struct {
	ID                  uint64    `json:"id,omitempty"`
	StartAt             time.Time `json:"startAt,omitempty"`
	FinalAt             time.Time `json:"finalAt,omitempty"`
	TotalWeightProducts int       `json:"totalWeightProducts,omitempty"`
	Man                 Man       `json:"man,omitempty"`
	Products            []Product `json:"products,omitempty"`
}

type Product struct {
	ID      uint64 `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Weight  int    `json:"weight,omitempty"`
	BarCode int    `json:"barCode,omitempty"`
}

type Man struct {
	ID          uint64      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	LastName    string      `json:"lastName,omitempty"`
	PhoneNumber string      `json:"phoneNumber,omitempty"`
	Transport   []Transport `json:"transport,omitempty"`
}

type Transport struct {
	ID   uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

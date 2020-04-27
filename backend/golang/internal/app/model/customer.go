package model

import "time"

// Customer ...
type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Phone     int       `json:"phone"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CustomerPhotos struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	Customer  Customer
	CreatedAt time.Time `json:"created_at"`
}

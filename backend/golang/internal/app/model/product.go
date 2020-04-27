package model

import "time"

// Product ...
type Product struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Category    *Category  `json:"category"`
	Stock       *Stock     `json:"stock"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

// ProductPhotos ...
type ProductPhotos struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	Product   Product
	CreatedAt *time.Time `json:"created_at"`
}

// Category - модель категорий
type Category struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// Stock ...
type Stock struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Precent     float64    `json:"precent"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

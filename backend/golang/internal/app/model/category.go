package model

import "time"

// Category - модель категорий
type Category struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreateAt    *time.Time `json:"create_at,omitempty"`
	UpdateAt    *time.Time `json:"update_at,omitempty"`
}

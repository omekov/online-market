package model

import "time"

// Category - модель категорий
type Category struct {
	ID       int        `json:"id"`
	Name     string     `json:"name"`
	RusName  string     `json:"rus_name"`
	Color    string     `json:"color"`
	CreateAt *time.Time `json:"create_at,omitempty"`
	UpdateAt *time.Time `json:"update_at,omitempty"`
	OriginID int        `json:"origin_id,"`
}

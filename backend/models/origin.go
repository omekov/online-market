package models

import (
	"time"

	"github.com/omekov/online-market/backend/db"
)

// FoodProduct ...
type FoodProduct struct {
	Origins []Origins `json:"origins,omitempty"`
}

// Origins ...
type Origins struct {
	ID          uint64      `json:"id,omitempty"`
	Name        string      `json:"name,omitempty"`
	RussianName string      `json:"russianName,omitempty` // RussianName ...
	CreateAt    time.Time   `json:"createAt,omitempty"`
	UpdateAt    time.Time   `json:"updateAt,omitempty"`
	Categories  *[]Category `json:"categories"`
}

// GetOrigins ...
func GetOrigins(product *FoodProduct) error {
	rows, err := db.DB.Query(`SELECT id, name, russian_name, updateAt, createAt FROM origins`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		org := Origins{}
		err = rows.Scan(
			&org.ID,
			&org.Name,
			&org.RussianName,
			&org.UpdateAt,
			&org.CreateAt,
		)
		if err != nil {
			return err
		}
		product.Origins = append(product.Origins, org)
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}

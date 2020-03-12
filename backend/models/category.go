package models

import (
	"time"

	"github.com/omekov/online-market/backend/db"
)

type Category struct {
	ID          int32      `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	RussianName string     `json:"russianName,omitempty` // RussianName ...
	Color       string     `json:"color,omitempty"`
	CreateAt    time.Time  `json:"createAt,omitempty"`
	UpdateAt    *time.Time `json:"updateAt"`
	OriginID    int32      `json:"originId,omitempty"`
}

// GetCategory ...
func GetCategories(categories *[]Category) error {
	rows, err := db.DB.Query(`
		SELECT 
		id, 
		name, 
		russianName, 
		color, 
		updateAt, 
		createAt, 
		originId 
		FROM categories 
		ORDER BY createAt DESC;`,
	)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		cat := Category{}
		err = rows.Scan(
			&cat.ID,
			&cat.Name,
			&cat.RussianName,
			&cat.Color,
			&cat.UpdateAt,
			&cat.CreateAt,
			&cat.OriginID,
		)
		if err != nil {
			return err
		}
		*categories = append(*categories, cat)
	}
	return nil
}

// GetCategory ...
func GetCategory(id int32, category *Category) error {
	err := db.DB.QueryRow(`
		SELECT 
		id, 
		name, 
		russianName, 
		color, 
		updateAt, 
		createAt 
		FROM categories 
		WHERE id = $1`,
		id,
	).Scan(
		&category.ID,
		&category.Name,
		&category.RussianName,
		&category.Color,
		&category.UpdateAt,
		&category.CreateAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// InsertCategory ...
func SaveCategory(category *Category) error {
	_, err := db.DB.Exec(`
		INSERT INTO categories (
		name,
		russianName,
		color,
		originId
		) VALUES ($1,$2,$3,$4);`,
		&category.Name,
		&category.RussianName,
		&category.Color,
		&category.OriginID,
	)
	if err != nil {
		return err
	}
	return err
}

// UpdateCategory ...
func UpdateCategory(id int32, category *Category) error {
	_, err := db.DB.Exec(`
		UPDATE categories
		SET name = $2,
		russianName = $3,
		color = $4,
		updateAt = $5,
		originId = $6
		WHERE id = $1;`,
		id,
		&category.Name,
		&category.RussianName,
		&category.Color,
		time.Now(),
		&category.OriginID,
	)
	if err != nil {
		return err
	}
	return nil
}

func DeleteCategory(id int32) error {
	_, err := db.DB.Exec(`
		DELETE FROM categories
		WHERE id = $1;`,
		id,
	)
	if err != nil {
		return err
	}
	return nil
}

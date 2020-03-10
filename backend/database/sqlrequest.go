package database

import (
	"fmt"
	"log"
	"time"

	"github.com/omekov/online-market/backend/models"
)

// SelectOrigin ...
func SelectOrigin(product *models.FoodProduct) error {
	rows, err := db.Query(`SELECT id, name, russian_name, update_at, create_at FROM origins`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		org := models.Origins{}
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

// InsertCategory ...
func InsertCategory() {
	insertCategory := `
	INSERT INTO category (
	name,
	russian_name,
	color,
	create_at,
	update_at,
	origin_id
	) VALUES ($1,$2,$3,$4,$5,$6);`
	_, err := db.Exec(insertCategory, "Mushrooms", "Грибы", "#C0C0C0", time.Now(), time.Now(), 1)
	if err != nil {
		log.Fatalf("insert Catergory", err)
	}
	fmt.Println("Success Insert")
}

// UpdateCategory ...
func UpdateCategory() {
	updateCategory := `
	UPDATE category
	SET russian_name = $2
	WHERE id = $1
	RETURNING id, name;`
	var name string
	var id int
	err := db.QueryRow(updateCategory, 4, "Грибок").Scan(&id, &name)
	if err != nil {
		log.Fatalf("update Catergory", err)
	}
	fmt.Println(id, name)
}

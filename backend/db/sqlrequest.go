package db

import (
	"fmt"
	"log"
	"time"

	"github.com/omekov/online-market/backend/models"
)

func GetOrigin(id int) (*models.Origins, error) {
	origin := models.Origins{}
	err := db.QueryRow(`SELECT id, name, russian_name, update_at, create_at FROM origins WHERE id = $1`, id).Scan(&origin.ID, &origin.Name, &origin.RussianName, &origin.UpdateAt, &origin.CreateAt)
	if err != nil {
		return nil, err
	}
	return &origin, nil
}
func GetAllOrigins(product *models.FoodProduct) error {
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
			&org.Categories,
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
func GetCategory(id int) (*models.Category, error) {
	category := models.Category{}
	err := db.QueryRow(`SELECT id, name, russian_name, color, update_at, create_at FROM category WHERE id = $1`, id).Scan(&category.ID, &category.Name, &category.RussianName, &category.Color, &category.UpdateAt, &category.CreateAt)
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// func GetCategoryByOriginId(id uint64) (*[]models.Category, error) {
// 	category := make([]models.Category, 1)
// 	rows, err := db.Query(`SELECT id, name, russian_name, color, update_at, create_at FROM category WHERE origin_id = $1`, id)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		cat := models.Category{}
// 		err = rows.Scan(
// 			&cat.ID,
// 			&cat.Name,
// 			&cat.RussianName,
// 			&cat.Color,
// 			&cat.UpdateAt,
// 			&cat.CreateAt,
// 		)
// 		if err != nil {
// 			return nil, err
// 		}
// 		product.Origns = append(product.Origins, org)
// 	}
// 	return &category, nil
// }
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

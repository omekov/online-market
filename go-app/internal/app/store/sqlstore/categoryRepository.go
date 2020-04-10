package sqlstore

import (
	"github.com/omekov/online-market/go-app/internal/app/model"
)

// CategoryRepository - Модель для Репозиторий
type CategoryRepository struct {
	store *Store
}

// GetAll - Возвращает все категорий
func (r *CategoryRepository) GetAll() ([]model.Category, error) {
	var categories []model.Category
	rows, err := r.store.db.Query("SELECT id, name, russianName, color, originId FROM categories ORDER BY createAt DESC;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		c := model.Category{}
		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.RusName,
			&c.Color,
			&c.OriginID,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	return categories, nil
}

// GetByID - Возвращает по ID категорий
func (r *CategoryRepository) GetByID(id int, category *model.Category) error {
	err := r.store.db.QueryRow(
		"SELECT id, name, russianName, color, updateAt, createAt, originId FROM categories WHERE id = $1;",
		id,
	).Scan(
		&category.ID,
		&category.Name,
		&category.RusName,
		&category.Color,
		&category.UpdateAt,
		&category.CreateAt,
		&category.OriginID,
	)
	if err != nil {
		return err
	}
	return nil
}

// Create - Создаем категорию
func (r *CategoryRepository) Create(category *model.Category) error {
	err := r.store.db.QueryRow(
		"INSERT INTO categories (name, russianName, color, originId) VALUES ($1,$2,$3,$4) RETURNING id;",
		&category.Name,
		&category.RusName,
		&category.Color,
		&category.OriginID,
	).Scan(&category.ID)
	if err != nil {
		return err
	}
	return nil
}

// Update - Обновляем категорию по ID
func (r *CategoryRepository) Update(id int, category *model.Category) error {
	err := r.store.db.QueryRow(
		"UPDATE categories SET name = $2, russianName = $3, color = $4, updateAt = $5,	originId = $6 WHERE id = $1 RETURNING id;",
		id,
		&category.Name,
		&category.RusName,
		&category.Color,
		&category.UpdateAt,
		&category.OriginID,
	).Scan(&category.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete - Удаляем категорию по ID
func (r *CategoryRepository) Delete(id int) error {
	err := r.store.db.QueryRow(
		"DELETE FROM categories WHERE id = $1;",
		id,
	).Scan(id)
	if err != nil {
		return err
	}
	return nil
}

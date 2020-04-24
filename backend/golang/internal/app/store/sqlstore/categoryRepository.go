package sqlstore

import (
	"github.com/omekov/online-market/backend/golang/internal/app/model"
)

// CategoryRepository - Модель для Репозиторий
type CategoryRepository struct {
	store *Store
}

// GetAll - Возвращает все категорий
func (r *CategoryRepository) GetAll() ([]model.Category, error) {
	var categories []model.Category
	rows, err := r.store.db.Query("SELECT id, name, description FROM categories ORDER BY created_at DESC;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		c := model.Category{}
		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.Description,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

// GetByID - Возвращает по ID категорий
func (r *CategoryRepository) GetByID(id int, category *model.Category) error {
	err := r.store.db.QueryRow(
		"SELECT id, name, description, updated_at, created_at FROM categories WHERE id = $1;",
		id,
	).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.UpdateAt,
		&category.CreateAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// Create - Создаем категорию
func (r *CategoryRepository) Create(category *model.Category) error {
	err := r.store.db.QueryRow(
		"INSERT INTO categories (name, description) VALUES ($1,$2) RETURNING id;",
		&category.Name,
		&category.Description,
	).Scan(&category.ID)
	if err != nil {
		return err
	}
	return nil
}

// Update - Обновляем категорию по ID
func (r *CategoryRepository) Update(id int, category *model.Category) error {
	err := r.store.db.QueryRow(
		"UPDATE categories SET name = $2, description = $3, updated_at = $4 WHERE id = $1 RETURNING id;",
		id,
		&category.Name,
		&category.Description,
		&category.UpdateAt,
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

package sqlstore

import (
	"database/sql"
	"time"

	"github.com/omekov/online-market/go-app/internal/app/model"
	"github.com/omekov/online-market/go-app/internal/app/store"
)

// CategoryRepository - Модель для Репозиторий
type CategoryRepository struct {
	store *Store
}

// GetAll - Возвращает все категорий
func (r *CategoryRepository) GetAll(categories *[]model.Category) error {
	rows, err := r.store.db.Query("SELECT id, name, russianName, color, createAt, updateAt, originId FROM categories ORDER BY createAt DESC;")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		c := model.Category{}
		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.RusName,
			&c.Color,
			&c.CreateAt,
			&c.UpdateAt,
			&c.OriginID,
		)
		if err != nil {
			return err
		}
		*categories = append(*categories, c)
	}
	return nil
}

// GetByID - Возвращает по ID категорий
func (r *CategoryRepository) GetByID(id int, category *model.Category) error {
	err := r.store.db.QueryRow(
		"SELECT id, name, russianName, color, updateAt, createAt FROM categories WHERE id = $1;",
		id,
	).Scan(
		&category.ID,
		&category.Name,
		&category.RusName,
		&category.Color,
		&category.UpdateAt,
		&category.CreateAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}
	return nil
}

// Create - Создаем категорию
func (r *CategoryRepository) Create(category *model.Category) error {
	_, err := r.store.db.Exec(
		"INSERT INTO categories (name, russianName, color, originId) VALUES ($1,$2,$3,$4);",
		&category.Name,
		&category.RusName,
		&category.Color,
		&category.OriginID,
	)
	if err != nil {
		return err
	}
	return nil
}

// Update - Обновляем категорию по ID
func (r *CategoryRepository) Update(id int, category *model.Category) error {
	_, err := r.store.db.Exec(
		"UPDATE categories SET name = $2, russianName = $3, color = $4, updateAt = $5,	originId = $6 WHERE id = $1;",
		id,
		&category.Name,
		&category.RusName,
		&category.Color,
		time.Now(),
		&category.OriginID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}
	return nil
}

// Delete - Удаляем категорию по ID
func (r *CategoryRepository) Delete(id int) error {
	_, err := r.store.db.Exec(
		"DELETE FROM categories WHERE id = $1;",
		id,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return store.ErrRecordNotFound
		}
		return err
	}
	return nil
}

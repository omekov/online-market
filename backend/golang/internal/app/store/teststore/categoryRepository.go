package teststore

import "github.com/omekov/online-market/backend/golang/internal/app/model"

// CategoryRepository - модель
type CategoryRepository struct {
	store      *Store
	categories map[int]*model.Category
}

// GetAll - проверка
func (r *CategoryRepository) GetAll() ([]model.Category, error) {
	return nil, nil
}

// GetByID - проверка
func (r *CategoryRepository) GetByID(int, *model.Category) error {
	return nil
}

// Create - проверка
func (r *CategoryRepository) Create(*model.Category) error {
	return nil
}

// Update - проверка
func (r *CategoryRepository) Update(int, *model.Category) error {
	return nil
}

// Delete - проверка
func (r *CategoryRepository) Delete(int) error {
	return nil
}

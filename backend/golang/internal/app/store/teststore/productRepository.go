package teststore

import "github.com/omekov/online-market/backend/golang/internal/app/model"

// ProductRepository - модель
type ProductRepository struct {
	store    *Store
	products map[int]*model.Product
}

// GetAll - проверка
func (r *ProductRepository) GetAll() ([]model.Product, error) {
	return nil, nil
}

// GetByID - проверка
func (r *ProductRepository) GetByID(int, *model.Product) error {
	return nil
}

// Create - проверка
func (r *ProductRepository) Create(*model.Product) error {
	return nil
}

// Update - проверка
func (r *ProductRepository) Update(int, *model.Product) error {
	return nil
}

// Delete - проверка
func (r *ProductRepository) Delete(int) error {
	return nil
}

// CreateStock - проверка
func (r *ProductRepository) CreateStock(*model.Stock) error {
	return nil
}

// CreateCategory - проверка
func (r *ProductRepository) CreateCategory(*model.Category) error {
	return nil
}

func (r *ProductRepository) GetAllCategory() ([]model.Category, error) {
	return nil, nil
}

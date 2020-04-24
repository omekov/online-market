package teststore

import "github.com/omekov/online-market/backend/golang/internal/app/model"

// CustomerRepository - модель
type CustomerRepository struct {
	store     *Store
	customers map[int]*model.Customer
}

// GetAll - проверка
func (r *CustomerRepository) GetAll() ([]model.Customer, error) {
	return nil, nil
}

// GetByID - проверка
func (r *CustomerRepository) GetByID(int, *model.Customer) error {
	return nil
}

// Create - проверка
func (r *CustomerRepository) Create(*model.Customer) error {
	return nil
}

// Update - проверка
func (r *CustomerRepository) Update(int, *model.Customer) error {
	return nil
}

// Delete - проверка
func (r *CustomerRepository) Delete(int) error {
	return nil
}

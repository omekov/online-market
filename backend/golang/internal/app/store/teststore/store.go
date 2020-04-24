package teststore

import (
	"github.com/omekov/online-market/backend/golang/internal/app/model"
	"github.com/omekov/online-market/backend/golang/internal/app/store"
)

// Store - это тестовый store для теста
type Store struct {
	categoryRepository *CategoryRepository
	productRepository  *ProductRepository
	customerRepository *CustomerRepository
}

// New -
func New() *Store {
	return &Store{}
}

// Category -
func (s *Store) Category() store.CategoryRepositorer {
	if s.categoryRepository != nil {
		return s.categoryRepository
	}
	s.categoryRepository = &CategoryRepository{
		store:      s,
		categories: make(map[int]*model.Category),
	}
	return s.categoryRepository

}

// Product - ...
func (s *Store) Product() store.ProductRepositorer {
	if s.productRepository != nil {
		return s.productRepository
	}
	s.productRepository = &ProductRepository{
		store:    s,
		products: make(map[int]*model.Product),
	}
	return s.productRepository
}

// Customer - ...
func (s *Store) Customer() store.CustomerRepositorer {
	if s.customerRepository != nil {
		return s.customerRepository
	}
	s.customerRepository = &CustomerRepository{
		store:     s,
		customers: make(map[int]*model.Customer),
	}
	return s.customerRepository
}

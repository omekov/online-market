package teststore

import (
	"github.com/omekov/online-market/backend/golang/internal/app/model"
	"github.com/omekov/online-market/backend/golang/internal/app/store"
)

// Store - это тестовый store для теста
type Store struct {
	productRepository  *ProductRepository
	customerRepository *CustomerRepository
	cartRepository     *CartRepository
}

// New -
func New() *Store {
	return &Store{}
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

func (s *Store) Cart() store.CartRepositorer {
	if s.cartRepository != nil {
		return s.cartRepository
	}
	s.cartRepository = &CartRepository{
		store: s,
		cart:  make(map[int]*model.Cart),
	}
	return s.cartRepository
}

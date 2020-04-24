package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" //postgres driver
	"github.com/omekov/online-market/backend/golang/internal/app/store"
)

// Store - объявления модели
type Store struct {
	db                 *sql.DB
	categoryRepository *CategoryRepository
	productRepository  *ProductRepository
	customerRepository *CustomerRepository
}

// New - метод для соединение базы
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Category - Главный метод для обращаения внешний мир
func (s *Store) Category() store.CategoryRepositorer {
	if s.categoryRepository != nil {
		return s.categoryRepository
	}
	s.categoryRepository = &CategoryRepository{
		store: s,
	}
	return s.categoryRepository
}

// Product - ...
func (s *Store) Product() store.ProductRepositorer {
	if s.productRepository != nil {
		return s.productRepository
	}
	s.productRepository = &ProductRepository{
		store: s,
	}
	return s.productRepository
}

// Customer - ...
func (s *Store) Customer() store.CustomerRepositorer {
	if s.customerRepository != nil {
		return s.customerRepository
	}
	s.customerRepository = &CustomerRepository{
		store: s,
	}
	return s.customerRepository
}

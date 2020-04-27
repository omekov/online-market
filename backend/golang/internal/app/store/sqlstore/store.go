package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" //postgres driver
	"github.com/omekov/online-market/backend/golang/internal/app/store"
)

// Store - объявления модели
type Store struct {
	db                 *sql.DB
	productRepository  *ProductRepository
	customerRepository *CustomerRepository
	cartRepository     *CartRepository
}

// New - метод для соединение базы
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
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

func (s *Store) Cart() store.CartRepositorer {
	if s.cartRepository != nil {
		return s.cartRepository
	}
	s.cartRepository = &CartRepository{
		store: s,
	}
	return s.cartRepository
}

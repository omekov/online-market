package sqlstore

import (
	"database/sql"

	_ "github.com/lib/pq" //postgres driver
	"github.com/omekov/online-market/back-api/internal/app/store"
)

// Store - объявления модели
type Store struct {
	db                 *sql.DB
	categoryRepository *CategoryRepository
}

// New - метод для соединение базы
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// Category - Главный метод для обращаения внешний мир
func (s *Store) Category() store.CategoryRepository {
	if s.categoryRepository != nil {
		return s.categoryRepository
	}
	s.categoryRepository = &CategoryRepository{
		store: s,
	}
	return s.categoryRepository
}

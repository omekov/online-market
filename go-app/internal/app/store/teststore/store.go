package teststore

import (
	"github.com/omekov/online-market/go-app/internal/app/model"
	"github.com/omekov/online-market/go-app/internal/app/store"
)

// Store - это тестовый store для теста
type Store struct {
	categoryRepository *CategoryRepository
}

// New -
func New() *Store {
	return &Store{}
}

// Category -
func (s *Store) Category() store.CategoryRepository {
	if s.categoryRepository != nil {
		return s.categoryRepository
	}
	s.categoryRepository = &CategoryRepository{
		store:      s,
		categories: make(map[int]*model.Category),
	}
	return s.categoryRepository

}

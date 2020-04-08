package store

import "github.com/omekov/online-market/go-app/internal/app/model"

// CategoryRepository - какие методы есть в категорий
type CategoryRepository interface {
	GetAll(*[]model.Category) error
	GetByID(int, *model.Category) error
	Create(*model.Category) error
	Update(int, *model.Category) error
	Delete(int) error
}

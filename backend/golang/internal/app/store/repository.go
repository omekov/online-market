package store

import "github.com/omekov/online-market/backend/golang/internal/app/model"

// CategoryRepositorer - какие методы есть в категорий
type CategoryRepositorer interface {
	GetAll() ([]model.Category, error)
	GetByID(int, *model.Category) error
	Create(*model.Category) error
	Update(int, *model.Category) error
	Delete(int) error
}

// ProductRepositorer - какие методы есть в продуктов
type ProductRepositorer interface {
	GetAll() ([]model.Product, error)
	GetByID(int, *model.Product) error
	Create(*model.Product) error
	Update(int, *model.Product) error
	Delete(int) error
}

// CustomerRepositorer - какие методы есть у клиентов
type CustomerRepositorer interface {
	GetAll() ([]model.Customer, error)
	GetByID(int, *model.Customer) error
	Create(*model.Customer) error
	Update(int, *model.Customer) error
	Delete(int) error
}

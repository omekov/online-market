package store

import "github.com/omekov/online-market/backend/golang/internal/app/model"

// ProductRepositorer - какие методы есть в продуктов
type ProductRepositorer interface {
	GetAll() ([]model.Product, error)
	GetByID(int, *model.Product) error
	Create(*model.Product) error
	Update(int, *model.Product) error
	Delete(int) error
	CreateCategory(*model.Category) error
	CreateStock(*model.Stock) error
}

// CustomerRepositorer - какие методы есть у клиентов
type CustomerRepositorer interface {
	GetAll() ([]model.Customer, error)
	GetByID(int, *model.Customer) error
	Create(*model.Customer) error
	Update(int, *model.Customer) error
	Delete(int) error
}

type CartRepositorer interface {
	GetByID(int, *model.Cart) error
	Create(*model.Cart) error
	Update(int, *model.Cart) error
	GetByCustomerID(int) ([]model.Cart, error)
	GetByCartIDandProductID(int, int) (*model.CartProduct, error)
}

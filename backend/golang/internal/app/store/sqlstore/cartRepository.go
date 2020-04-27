package sqlstore

import "github.com/omekov/online-market/backend/golang/internal/app/model"

type CartRepository struct {
	store *Store
}

func (r *CartRepository) GetByID(id int, cart *model.Cart) error {
	return nil
}

func (r *CartRepository) Create(cart *model.Cart) error {
	return nil
}

func (r *CartRepository) Update(id int, cart *model.Cart) error {
	return nil
}

func (r *CartRepository) GetByCustomerID(customerID int) ([]model.Cart, error) {
	return nil, nil
}

func (r *CartRepository) GetByCartIDandProductID(cartID int, customerID int) (*model.CartProduct, error) {
	return nil, nil
}

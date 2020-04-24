package sqlstore

import "github.com/omekov/online-market/backend/golang/internal/app/model"

// CustomerRepository ...
type CustomerRepository struct {
	store *Store
}

// GetAll ...
func (r *CustomerRepository) GetAll() ([]model.Customer, error) {
	var customers []model.Customer
	rows, err := r.store.db.Query("SELECT id, name, lastname, phone, email, created_at, updated_at FROM customers ORDER BY created_at DESC;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		c := model.Customer{}
		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.Lastname,
			&c.Phone,
			&c.Email,
			&c.CreatedAt,
			&c.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return customers, nil
}

// GetByID ...
func (r *CustomerRepository) GetByID(id int, customer *model.Customer) error {
	err := r.store.db.QueryRow("SELECT id, name, lastname, phone, email, created_at, updated_at FROM customers WHERE id = $1;",
		id,
	).Scan(
		&customer.Name,
		&customer.Lastname,
		&customer.Phone,
		&customer.Email,
		&customer.CreatedAt,
		&customer.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// Create ...
func (r *CustomerRepository) Create(customer *model.Customer) error {
	err := r.store.db.QueryRow("INSERT INTO customers (name, lastname, phone, email, updated_at) VALUES($1, $2, $3, $4, $5) RETURNING id;",
		&customer.Name,
		&customer.Lastname,
		&customer.Phone,
		&customer.Email,
		&customer.UpdatedAt,
	).Scan(&customer.ID)
	if err != nil {
		return err
	}
	return nil
}

// Update ...
func (r *CustomerRepository) Update(id int, customer *model.Customer) error {
	err := r.store.db.QueryRow(
		"UPDATE customers SET name = $2, lastname = $3, phone = $4, email = $5, updated_at = $6 WHERE id = $1 RETURNING id",
		id,
		&customer.Name,
		&customer.Lastname,
		&customer.Phone,
		&customer.Email,
		&customer.UpdatedAt,
	).Scan(&customer.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete ...
func (r *CustomerRepository) Delete(id int) error {
	err := r.store.db.QueryRow(
		"DELETE FROM customers WHERE id = $1",
		id,
	).Scan(id)
	if err != nil {
		return err
	}
	return nil
}

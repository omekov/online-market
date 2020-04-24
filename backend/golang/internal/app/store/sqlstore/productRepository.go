package sqlstore

import "github.com/omekov/online-market/backend/golang/internal/app/model"

// ProductRepository ...
type ProductRepository struct {
	store *Store
}

// GetAll ...
func (r *ProductRepository) GetAll() ([]model.Product, error) {
	products := make([]model.Product, 0)
	rows, err := r.store.db.Query("SELECT id, name, description, price, category_id, stock_id, created_at, updated_at FROM products ORDER BY createAt DESC;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := new(model.Product)
		err = rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.CategoryID,
			&p.StockID,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

// GetByID ...
func (r *ProductRepository) GetByID(id int, product *model.Product) error {
	err := r.store.db.QueryRow(
		"SELECT id, name, description, price, category_id, stock_id, created_at, updated_at FROM products WHERE id = $1;",
		id,
	).Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CategoryID,
		&product.StockID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// Create ...
func (r *ProductRepository) Create(product *model.Product) error {
	err := r.store.db.QueryRow(
		"INSERT INTO products (name, description, price, category_id, stock_id) VALUES($1,$2,$3,$4,$5) RETURNING id;",
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CategoryID,
		&product.StockID,
	).Scan(&product.ID)
	if err != nil {
		return err
	}
	// драйвер pq не поддерживает метод LastInsertId()
	// row, err := res.LastInsertId()
	// if err != nil {
	// 	return err
	// }
	// product.ID = int(row)
	return nil
}

// Update ...
func (r *ProductRepository) Update(id int, product *model.Product) error {
	err := r.store.db.QueryRow(
		"UPDATE products SET name = $2, description = $3, price = $4, category_id = $5, stock_id = $6, updated_at = $7 WHERE id = $1 RETURNING id;",
		id,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.CategoryID,
		&product.StockID,
		&product.UpdatedAt,
	).Scan(&product.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete ...
func (r *ProductRepository) Delete(id int) error {
	err := r.store.db.QueryRow(
		"DELETE FROM products WHERE id = $1;",
		id,
	).Scan(id)
	if err != nil {
		return err
	}
	return nil
}

package sqlstore

import (
	"database/sql"

	"github.com/omekov/online-market/backend/golang/internal/app/model"
	"github.com/omekov/online-market/backend/golang/internal/app/store"
)

// ProductRepository ...
type ProductRepository struct {
	store *Store
}

// GetAll ...
func (r *ProductRepository) GetAll() ([]model.Product, error) {
	products := make([]model.Product, 0)
	rows, err := r.store.db.Query("SELECT id, name, description, price, category_id, stock_id, created_at, updated_at FROM products ORDER BY created_at DESC;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		p := model.Product{Category: &model.Category{}, Stock: &model.Stock{}}
		err = rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.Category.ID,
			&p.Stock.ID,
			&p.CreatedAt,
			&p.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		if err != r.GetByCategoryID(p.Category.ID, p.Category) {
			return nil, err
		}
		if err != r.GetByStockID(p.Stock.ID, p.Stock) {
			return nil, err
		}
		products = append(products, p)
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
		&product.Category.ID,
		&product.Stock.ID,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return err
	}
	if err != r.GetByCategoryID(product.Category.ID, product.Category) {
		return err
	}
	if err != r.GetByStockID(product.Stock.ID, product.Stock) {
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
		&product.Category.ID,
		&product.Stock.ID,
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
		&product.Category.ID,
		&product.Stock.ID,
		&product.UpdatedAt,
	).Scan(&product.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete ...
func (r *ProductRepository) Delete(id int) error {
	restultRow, err := r.store.db.Exec(
		"DELETE FROM products WHERE id = $1;",
		id,
	)
	if err != nil {
		return err
	}
	row, err := restultRow.RowsAffected()
	if row == 0 {
		return sql.ErrNoRows
	}
	if err != nil {
		return err
	}
	return nil
}

// CreateStock ...
func (r *ProductRepository) CreateStock(stock *model.Stock) error {
	err := r.store.db.QueryRow(
		"INSERT INTO stocks (name, description, precent) VALUES($1, $2, $3) RETURNING id",
		&stock.Name,
		&stock.Description,
		&stock.Precent,
	).Scan(&stock.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetByStockID ...
func (r *ProductRepository) GetByStockID(id int, stock *model.Stock) error {
	err := r.store.db.QueryRow(
		"SELECT id, name, description, created_at, updated_at FROM stocks WHERE id = $1",
		id,
	).Scan(
		&stock.ID,
		&stock.Name,
		&stock.Description,
		&stock.CreatedAt,
		&stock.UpdatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetAllCategory - Возвращает все категорий
func (r *ProductRepository) GetAllCategory() ([]model.Category, error) {
	var categories []model.Category
	rows, err := r.store.db.Query("SELECT id, name, description FROM categories ORDER BY created_at DESC;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		c := model.Category{}
		err = rows.Scan(
			&c.ID,
			&c.Name,
			&c.Description,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

// GetByCategoryID - Возвращает по ID категорий
func (r *ProductRepository) GetByCategoryID(id int, category *model.Category) error {
	err := r.store.db.QueryRow(
		"SELECT id, name, description, updated_at, created_at FROM categories WHERE id = $1;",
		id,
	).Scan(
		&category.ID,
		&category.Name,
		&category.Description,
		&category.UpdatedAt,
		&category.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

// CreateCategory - Создаем категорию
func (r *ProductRepository) CreateCategory(category *model.Category) error {
	err := r.store.db.QueryRow(
		"INSERT INTO categories (name, description) VALUES ($1,$2) RETURNING id;",
		&category.Name,
		&category.Description,
	).Scan(&category.ID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateCategory - Обновляем категорию по ID
func (r *ProductRepository) UpdateCategory(id int, category *model.Category) error {
	err := r.store.db.QueryRow(
		"UPDATE categories SET name = $2, description = $3, updated_at = $4 WHERE id = $1 RETURNING id;",
		id,
		&category.Name,
		&category.Description,
		&category.UpdatedAt,
	).Scan(&category.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCategory - Удаляем категорию по ID
func (r *ProductRepository) DeleteCategory(id int) error {
	restultRow, err := r.store.db.Exec(
		"DELETE FROM categories WHERE id = $1;",
		id,
	)
	if row, err := restultRow.RowsAffected(); err != nil {
		if row == 0 {
			return store.ErrRecordNotFound
		}
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

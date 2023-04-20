package database

import (
	"database/sql"
	"fmt"

	"github.com/vitorconti/go-products/internal/entity"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{Db: db}
}

func (r *ProductRepository) Save(product *entity.Product) error {
	stmt, err := r.Db.Prepare("INSERT INTO products (name,description,price) VALUES (?, ?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.Name, product.Description, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) Edit(product *entity.Product) error {
	stmt, err := r.Db.Prepare("UPDATE products SET name=?, description=?, price=? WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.Name, product.Description, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}
func (r *ProductRepository) Find(limit, offset int) ([]entity.Product, error) {
	fmt.Sprintf("SELECT * FROM products ORDER BY 1 LIMIT %d OFFSET %d", limit, offset)
	rows, err := r.Db.Query(fmt.Sprintf("SELECT * FROM products ORDER BY 1 LIMIT %d OFFSET %d", limit, offset))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]entity.Product, 0)
	for rows.Next() {
		var product entity.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) FindOne(id int) (entity.Product, error) {
	var product entity.Product
	err := r.Db.QueryRow("SELECT * FROM products WHERE id=?", id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Product{}, fmt.Errorf("product with ID %d not found", id)
		}
		return entity.Product{}, err
	}

	return product, nil
}

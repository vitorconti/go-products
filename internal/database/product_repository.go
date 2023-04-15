package database

import (
	"database/sql"

	"github.com/vitorconti/go-products/internal/entity"
)

type ProductRepository struct {
	Db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{Db: db}
}

func (r *ProductRepository) Save(product *entity.Product) error {
	stmt, err := r.Db.Prepare("INSERT INTO products (name,description) VALUES (?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(product.Description, product.Price)
	if err != nil {
		return err
	}
	return nil
}

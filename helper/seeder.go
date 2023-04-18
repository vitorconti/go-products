package helper

import (
	"database/sql"

	"github.com/bxcodec/faker/v3"
)

func GenerateTableProduct(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(255),
			description VARCHAR(255),
			price DECIMAL(10,2)
		)
	`)
	if err != nil {
		panic(err)
	}
}
func SeedTableProduct(db *sql.DB) error {

	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM products")
	err := row.Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	for i := 0; i < 500; i++ {
		var product product
		err := faker.FakeData(&product)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec("INSERT INTO products (name,description, price) VALUES (?, ?, ?)", product.Name, product.Description, product.Price)
		if err != nil {
			return err
		}
	}

	return nil
}

type product struct {
	Name        string  `faker:"name"`
	Description string  `faker:"sentence"`
	Price       float32 `faker:"oneof: 4.95, 9.99, 31997.97"`
}

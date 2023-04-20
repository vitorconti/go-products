package entity

type ProductRepositoryInterface interface {
	Save(product *Product) (int64, error)
	Find(limit, offset int) ([]Product, error)
	FindOne(id int64) (Product, error)
	Remove(id int) error
	Edit(product *Product) error
}

package entity

type ProductRepositoryInterface interface {
	Save(product *Product) error
	Find(limit, offset int) ([]Product, error)
	FindOne(id int) (Product, error)
	Remove(id int) error
	Edit(product *Product) error
}

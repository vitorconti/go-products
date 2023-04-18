package entity

type ProductRepositoryInterface interface {
	Save(product *Product) error
	Find(limit, offset int) ([]Product, error)
}

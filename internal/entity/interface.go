package entity

type ProductRepositoryInterface interface {
	Save(product *Product) error
}

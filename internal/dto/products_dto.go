package dto

type ProductInputDTO struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type ProductOutputDTO struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
type ProductDeletdDTO struct {
	Message string
}
type ProductPaginationQueryParamsDTO struct {
	Page   int
	Limit  int
	Offset int
}

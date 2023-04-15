package usecase

import (
	"github.com/vitorconti/go-products/internal/entity"
	"github.com/vitorconti/go-products/pkg/events"
)

type ProductInputDTO struct {
	ID          string  `json:"id"`
	Description string  `json:"id"`
	Price       float64 `json:"price"`
}

type ProductOutputDTO struct {
	ID          string  `json:"id"`
	Description string  `json:"id"`
	Price       float64 `json:"price"`
}

type CreateProductUseCase struct {
	ProductRepository entity.ProductRepositoryInterface
	ProductCreated    events.EventInterface
	EventDispatcher   events.EventDispatcherInterface
}

func NewCreateProductUseCase(
	ProductRepository entity.ProductRepositoryInterface,
	ProductCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateProductUseCase {
	return &CreateProductUseCase{
		ProductRepository: ProductRepository,
		ProductCreated:    ProductCreated,
		EventDispatcher:   EventDispatcher,
	}
}

func (c *CreateProductUseCase) Execute(input ProductInputDTO) (ProductOutputDTO, error) {
	product := entity.Product{
		ID:          input.ID,
		Price:       input.Price,
		Description: input.Description,
	}

	if err := c.ProductRepository.Save(&product); err != nil {
		return ProductOutputDTO{}, err
	}

	outputDto := ProductOutputDTO{
		ID:          input.ID,
		Price:       input.Price,
		Description: input.Description,
	}

	c.ProductCreated.SetPayload(outputDto)
	c.EventDispatcher.Dispatch(c.ProductCreated)

	return outputDto, nil
}

package usecase

import (
	"github.com/vitorconti/go-products/internal/dto"
	"github.com/vitorconti/go-products/internal/entity"
	"github.com/vitorconti/go-products/pkg/events"
)

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

func (c *CreateProductUseCase) Execute(input dto.ProductInputDTO) (dto.ProductOutputDTO, error) {
	product := entity.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}

	if err := c.ProductRepository.Save(&product); err != nil {
		return dto.ProductOutputDTO{}, err
	}

	outputDto := dto.ProductOutputDTO{
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
	}

	c.ProductCreated.SetPayload(outputDto)
	c.EventDispatcher.Dispatch(c.ProductCreated)

	return outputDto, nil
}

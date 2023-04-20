package usecase

import (
	"github.com/vitorconti/go-products/internal/dto"
	"github.com/vitorconti/go-products/internal/entity"
	"github.com/vitorconti/go-products/pkg/events"
)

type UpdateProductUseCase struct {
	ProductRepository entity.ProductRepositoryInterface
	ProductUpdated    events.EventInterface
	EventDispatcher   events.EventDispatcherInterface
}

func NewUpdateProductUseCase(
	ProductRepository entity.ProductRepositoryInterface,
	ProductUpdated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *UpdateProductUseCase {
	return &UpdateProductUseCase{
		ProductRepository: ProductRepository,
		ProductUpdated:    ProductUpdated,
		EventDispatcher:   EventDispatcher,
	}
}

func (c *UpdateProductUseCase) Execute(input dto.ProductInputDTO) (dto.ProductOutputDTO, error) {
	product := entity.Product{
		ID:          input.ID,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
	}
	if _, err := c.ProductRepository.FindOne(product.ID); err != nil {
		return dto.ProductOutputDTO{}, err
	}

	if err := c.ProductRepository.Edit(&product); err != nil {
		return dto.ProductOutputDTO{}, err
	}

	outputDto := dto.ProductOutputDTO{
		ID:          input.ID,
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
	}

	c.ProductUpdated.SetPayload(outputDto)
	c.EventDispatcher.Dispatch(c.ProductUpdated)

	return outputDto, nil
}

package usecase

import (
	"github.com/vitorconti/go-products/internal/dto"
	"github.com/vitorconti/go-products/internal/entity"
	"github.com/vitorconti/go-products/pkg/events"
)

type DeleteProductUseCase struct {
	ProductRepository entity.ProductRepositoryInterface
	ProductDeleted    events.EventInterface
	EventDispatcher   events.EventDispatcherInterface
}

func NewDeleteProductUseCase(
	ProductRepository entity.ProductRepositoryInterface,
	ProductDeleted events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *DeleteProductUseCase {
	return &DeleteProductUseCase{
		ProductRepository: ProductRepository,
		ProductDeleted:    ProductDeleted,
		EventDispatcher:   EventDispatcher,
	}
}

func (c *DeleteProductUseCase) Execute(id int) (dto.ProductDeletdDTO, error) {

	if err := c.ProductRepository.Remove(id); err != nil {
		return dto.ProductDeletdDTO{}, err
	}
	outputDto := dto.ProductDeletdDTO{
		Message: "Product deleted",
	}
	c.ProductDeleted.SetPayload(outputDto)
	c.EventDispatcher.Dispatch(c.ProductDeleted)

	return outputDto, nil
}

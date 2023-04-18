package usecase

import (
	"strconv"

	"github.com/vitorconti/go-products/internal/entity"
	"github.com/vitorconti/go-products/pkg/events"
)

type ProductPaginationQueryParamsDTO struct {
	Page   int
	Limit  int
	Offset int
}

type RetriveProductUseCase struct {
	ProductRepository entity.ProductRepositoryInterface
	ProductRetrived   events.EventInterface
	EventDispatcher   events.EventDispatcherInterface
}

func NewRetriveProductUseCase(
	ProductRepository entity.ProductRepositoryInterface,
	ProductRetrived events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *RetriveProductUseCase {
	return &RetriveProductUseCase{
		ProductRepository: ProductRepository,
		ProductRetrived:   ProductRetrived,
		EventDispatcher:   EventDispatcher,
	}
}

func (c *RetriveProductUseCase) Execute(input ProductPaginationQueryParamsDTO) ([]ProductOutputDTO, error) {

	retrievedProducts, err := c.ProductRepository.Find(input.Limit, input.Offset)
	if err != nil {
		return []ProductOutputDTO{}, err
	}
	outputDto := make([]ProductOutputDTO, 0, len(retrievedProducts))
	for _, retrievedProduct := range retrievedProducts {
		parseId := strconv.Itoa(retrievedProduct.ID)
		if err != nil {
			return nil, err
		}
		product := ProductOutputDTO{
			ID:          parseId,
			Price:       retrievedProduct.Price,
			Description: retrievedProduct.Description,
		}
		outputDto = append(outputDto, product)
	}

	c.ProductRetrived.SetPayload(outputDto)
	c.EventDispatcher.Dispatch(c.ProductRetrived)

	return outputDto, nil
}

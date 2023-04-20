package web

import (
	"github.com/labstack/echo/v4"
	"github.com/vitorconti/go-products/internal/dto"
	"github.com/vitorconti/go-products/internal/entity"
	"github.com/vitorconti/go-products/internal/usecase"
	"github.com/vitorconti/go-products/pkg/events"
	"net/http"
)

type CreateProductHandler struct {
	EventDispatcher     events.EventDispatcherInterface
	ProductRepository   entity.ProductRepositoryInterface
	ProductCreatedEvent events.EventInterface
}

func NewCreateProductHandler(
	EventDispatcher events.EventDispatcherInterface,
	ProductRepository entity.ProductRepositoryInterface,
	ProductCreatedEvent events.EventInterface,
) *CreateProductHandler {
	return &CreateProductHandler{
		EventDispatcher:     EventDispatcher,
		ProductRepository:   ProductRepository,
		ProductCreatedEvent: ProductCreatedEvent,
	}
}

func (h *CreateProductHandler) Create(c echo.Context) error {
	var dto dto.ProductInputDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "One more parameters could be wrong."})
	}

	createProduct := usecase.NewCreateProductUseCase(h.ProductRepository, h.ProductCreatedEvent, h.EventDispatcher)
	output, err := createProduct.Execute(dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, output)
}

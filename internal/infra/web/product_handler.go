package web

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vitorconti/go-products/internal/entity"
	"github.com/vitorconti/go-products/internal/usecase"
	"github.com/vitorconti/go-products/pkg/events"
)

type ProductHandler struct {
	EventDispatcher     events.EventDispatcherInterface
	ProductRepository   entity.ProductRepositoryInterface
	ProductCreatedEvent events.EventInterface
}

func NewProductHandler(
	EventDispatcher events.EventDispatcherInterface,
	ProductRepository entity.ProductRepositoryInterface,
	ProductCreatedEvent events.EventInterface,
) *ProductHandler {
	return &ProductHandler{
		EventDispatcher:     EventDispatcher,
		ProductRepository:   ProductRepository,
		ProductCreatedEvent: ProductCreatedEvent,
	}
}

func (h *ProductHandler) Create(c echo.Context) error {
	var dto usecase.ProductInputDTO
	err := json.NewDecoder(c.Request().Body).Decode(&dto)
	if err != nil {
		http.Error(c.Response().Writer, err.Error(), http.StatusBadRequest)
		return err
	}

	createProduct := usecase.NewCreateProductUseCase(h.ProductRepository, h.ProductCreatedEvent, h.EventDispatcher)
	output, err := createProduct.Execute(dto)
	if err != nil {
		http.Error(c.Response().Writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	err = json.NewEncoder(c.Response().Writer).Encode(output)
	if err != nil {
		http.Error(c.Response().Writer, err.Error(), http.StatusInternalServerError)
		return err
	}
	return nil
}

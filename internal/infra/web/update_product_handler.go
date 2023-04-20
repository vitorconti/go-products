package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vitorconti/go-products/internal/dto"
	"github.com/vitorconti/go-products/internal/entity"
	"github.com/vitorconti/go-products/internal/usecase"
	"github.com/vitorconti/go-products/pkg/events"
)

type UpdateProductHandler struct {
	EventDispatcher     events.EventDispatcherInterface
	ProductRepository   entity.ProductRepositoryInterface
	ProductUpdatedEvent events.EventInterface
}

func NewUpdateProductHandler(
	EventDispatcher events.EventDispatcherInterface,
	ProductRepository entity.ProductRepositoryInterface,
	ProductUpdatedEvent events.EventInterface,
) *UpdateProductHandler {
	return &UpdateProductHandler{
		EventDispatcher:     EventDispatcher,
		ProductRepository:   ProductRepository,
		ProductUpdatedEvent: ProductUpdatedEvent,
	}
}

func (h *UpdateProductHandler) Edit(c echo.Context) error {
	var dto dto.ProductInputDTO
	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "One more parameters could be wrong."})
	}

	updateProduct := usecase.NewUpdateProductUseCase(h.ProductRepository, h.ProductUpdatedEvent, h.EventDispatcher)
	output, err := updateProduct.Execute(dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, output)
}

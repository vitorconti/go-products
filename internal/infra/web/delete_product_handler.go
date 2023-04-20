package web

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vitorconti/go-products/internal/entity"
	"github.com/vitorconti/go-products/internal/usecase"
	"github.com/vitorconti/go-products/pkg/events"
)

type DeleteProductHandler struct {
	EventDispatcher     events.EventDispatcherInterface
	ProductRepository   entity.ProductRepositoryInterface
	ProductDeletedEvent events.EventInterface
}

func NewDeleteProductHandler(
	EventDispatcher events.EventDispatcherInterface,
	ProductRepository entity.ProductRepositoryInterface,
	ProductDeletedEvent events.EventInterface,
) *DeleteProductHandler {
	return &DeleteProductHandler{
		EventDispatcher:     EventDispatcher,
		ProductRepository:   ProductRepository,
		ProductDeletedEvent: ProductDeletedEvent,
	}
}

func (h *DeleteProductHandler) Remove(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "One more parameters could be wrong."})
	}

	deleteProduct := usecase.NewDeleteProductUseCase(h.ProductRepository, h.ProductDeletedEvent, h.EventDispatcher)
	output, err := deleteProduct.Execute(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, output)
}

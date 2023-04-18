package web

import (
	"net/http"
	"strconv"

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

func (h *ProductHandler) Retrive(c echo.Context) error {
	var dto usecase.ProductPaginationQueryParamsDTO

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}
	dto.Page = page

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 10
	}
	dto.Limit = limit

	offset := (page) * limit
	dto.Offset = offset
	retrievedProduct := usecase.NewRetriveProductUseCase(h.ProductRepository, h.ProductCreatedEvent, h.EventDispatcher)
	output, err := retrievedProduct.Execute(dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, output)

}

package web

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vitorconti/go-products/internal/dto"
	"github.com/vitorconti/go-products/internal/entity"
	"github.com/vitorconti/go-products/internal/usecase"
	"github.com/vitorconti/go-products/pkg/events"
)

type RetriveProductHandler struct {
	EventDispatcher      events.EventDispatcherInterface
	ProductRepository    entity.ProductRepositoryInterface
	ProductRetrivedEvent events.EventInterface
}

func NewRetriveProductHandler(
	EventDispatcher events.EventDispatcherInterface,
	ProductRepository entity.ProductRepositoryInterface,
	ProductRetrivedEvent events.EventInterface,
) *RetriveProductHandler {
	return &RetriveProductHandler{
		EventDispatcher:      EventDispatcher,
		ProductRepository:    ProductRepository,
		ProductRetrivedEvent: ProductRetrivedEvent,
	}
}

func (h *RetriveProductHandler) Retrive(c echo.Context) error {
	var dto dto.ProductPaginationQueryParamsDTO

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

	offset := (page - 1) * limit
	dto.Offset = offset
	retrievedProduct := usecase.NewRetriveProductUseCase(h.ProductRepository, h.ProductRetrivedEvent, h.EventDispatcher)
	output, err := retrievedProduct.Execute(dto)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, output)

}

package web

import (
	"encoding/json"
	"net/http"

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

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.ProductInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createProduct := usecase.NewCreateProductUseCase(h.ProductRepository, h.ProductCreatedEvent, h.EventDispatcher)
	output, err := createProduct.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

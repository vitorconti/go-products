//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/vitorconti/go-products/internal/database"
	"github.com/vitorconti/go-products/internal/entity"
	"github.com/vitorconti/go-products/internal/event"
	"github.com/vitorconti/go-products/internal/infra/web"
	"github.com/vitorconti/go-products/internal/usecase"
	"github.com/vitorconti/go-products/pkg/events"
)

var setProductRepositoryDependency = wire.NewSet(
	database.NewProductRepository,
	wire.Bind(new(entity.ProductRepositoryInterface), new(*database.ProductRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewProductCreated,
	wire.Bind(new(events.EventInterface), new(*event.ProductCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setProductCreatedEvent = wire.NewSet(
	event.NewProductCreated,
	wire.Bind(new(events.EventInterface), new(*event.ProductCreated)),
)

func NewCreateProductUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateProductUseCase {
	wire.Build(
		setProductRepositoryDependency,
		setProductCreatedEvent,
		usecase.NewCreateProductUseCase,
	)
	return &usecase.CreateProductUseCase{}
}

func ProductHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.ProductHandler {
	wire.Build(
		setProductRepositoryDependency,
		setProductCreatedEvent,
		web.NewProductHandler,
	)
	return &web.ProductHandler{}
}

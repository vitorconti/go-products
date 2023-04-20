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
	event.NewProductUpdated,
	event.NewProductRetrived,
	wire.Bind(new(events.EventInterface), new(*event.ProductCreated)),
	wire.Bind(new(events.EventInterface), new(*event.ProductUpdated)),
	wire.Bind(new(events.EventInterface), new(*event.ProductRetrived)),
	wire.Bind(new(events.EventInterface), new(*event.ProductDeleted)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setProductCreatedEvent = wire.NewSet(
	event.NewProductCreated,
	wire.Bind(new(events.EventInterface), new(*event.ProductCreated)),
)
var setProductUpdatedEvent = wire.NewSet(
	event.NewProductUpdated,
	wire.Bind(new(events.EventInterface), new(*event.ProductUpdated)),
)
var setProductRetrivedEvent = wire.NewSet(
	event.NewProductRetrived,
	wire.Bind(new(events.EventInterface), new(*event.ProductRetrived)),
)
var setProductDeletedEvent = wire.NewSet(
	event.NewProductDeleted,
	wire.Bind(new(events.EventInterface), new(*event.ProductDeleted)),
)

func NewCreateProductUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateProductUseCase {
	wire.Build(
		setProductRepositoryDependency,
		setProductCreatedEvent,
		usecase.NewCreateProductUseCase,
	)
	return &usecase.CreateProductUseCase{}
}

func NewRetriveProductUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.RetriveProductUseCase {
	wire.Build(
		setProductRepositoryDependency,
		setProductRetrivedEvent,
		usecase.NewRetriveProductUseCase,
	)
	return &usecase.RetriveProductUseCase{}
}
func NewUpdateProductUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.UpdateProductUseCase {
	wire.Build(
		setProductRepositoryDependency,
		setProductUpdatedEvent,
		usecase.NewUpdateProductUseCase,
	)
	return &usecase.UpdateProductUseCase{}
}

func NewDeleteProductUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.DeleteProductUseCase {
	wire.Build(
		setProductRepositoryDependency,
		setProductDeletedEvent,
		usecase.NewDeleteProductUseCase,
	)
	return &usecase.DeleteProductUseCase{}
}

func CreateProductHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.CreateProductHandler {
	wire.Build(
		setProductRepositoryDependency,
		setProductCreatedEvent,
		web.NewCreateProductHandler,
	)
	return &web.CreateProductHandler{}
}

func RetriveProductHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.RetriveProductHandler {
	wire.Build(
		setProductRepositoryDependency,
		setProductRetrivedEvent,
		web.NewRetriveProductHandler,
	)
	return &web.RetriveProductHandler{}
}

func UpdateProductHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.UpdateProductHandler {
	wire.Build(
		setProductRepositoryDependency,
		setProductUpdatedEvent,
		web.NewUpdateProductHandler,
	)
	return &web.UpdateProductHandler{}
}
func DeleteProductHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.DeleteProductHandler {
	wire.Build(
		setProductRepositoryDependency,
		setProductDeletedEvent,
		web.NewDeleteProductHandler,
	)
	return &web.DeleteProductHandler{}
}

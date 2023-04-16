// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

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

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewCreateProductUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateProductUseCase {
	productRepository := database.NewProductRepository(db)
	productCreated := event.NewProductCreated()
	createProductUseCase := usecase.NewCreateProductUseCase(productRepository, productCreated, eventDispatcher)
	return createProductUseCase
}

func ProductHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.ProductHandler {
	productRepository := database.NewProductRepository(db)
	productCreated := event.NewProductCreated()
	productHandler := web.NewProductHandler(eventDispatcher, productRepository, productCreated)
	return productHandler
}

// wire.go:

var setProductRepositoryDependency = wire.NewSet(database.NewProductRepository, wire.Bind(new(entity.ProductRepositoryInterface), new(*database.ProductRepository)))

var setEventDispatcherDependency = wire.NewSet(events.NewEventDispatcher, event.NewProductCreated, wire.Bind(new(events.EventInterface), new(*event.ProductCreated)), wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)))

var setProductCreatedEvent = wire.NewSet(event.NewProductCreated, wire.Bind(new(events.EventInterface), new(*event.ProductCreated)))
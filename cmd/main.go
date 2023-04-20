package main

import (
	"database/sql"
	"fmt"

	// mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
	"github.com/vitorconti/go-products/configs"
	"github.com/vitorconti/go-products/helper"
	"github.com/vitorconti/go-products/internal/event/handler"
	"github.com/vitorconti/go-products/internal/infra/web/webserver"
	"github.com/vitorconti/go-products/pkg/events"
)

func main() {
	loadedConfigs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(loadedConfigs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", loadedConfigs.DBUser, loadedConfigs.DBPassword, loadedConfigs.DBHost, loadedConfigs.DBPort, loadedConfigs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	helper.GenerateTableProduct(db)
	helper.SeedTableProduct(db)
	rabbitMQChannel := createRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("ProductCreated", &handler.ProductCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	eventDispatcher.Register("ProductRetrived", &handler.ProductRetrivedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})
	eventDispatcher.Register("ProductUpdated", &handler.ProductUpdatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	webserver := webserver.NewWebServer(loadedConfigs.WebServerPort)
	createProductHandler := CreateProductHandler(db, eventDispatcher)
	retriveProductHandler := RetriveProductHandler(db, eventDispatcher)
	updateProductHandler := UpdateProductHandler(db, eventDispatcher)
	deleteProductHandler := DeleteProductHandler(db, eventDispatcher)
	webserver.AddHandler("POST", "/product", createProductHandler.Create)
	webserver.AddHandler("GET", "/product", retriveProductHandler.Retrive)
	webserver.AddHandler("PATCH", "/product", updateProductHandler.Edit)
	webserver.AddHandler("DELETE", "/product/:id", deleteProductHandler.Remove)
	fmt.Println("Starting web server on port", loadedConfigs.WebServerPort)
	webserver.Start()

}
func createRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}

package main

import (
	"database/sql"
	"fmt"

	// mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/streadway/amqp"
	"github.com/vitorconti/go-products/configs"
	"github.com/vitorconti/go-products/internal/event/handler"
	"github.com/vitorconti/go-products/internal/infra/web/webserver"
	"github.com/vitorconti/go-products/pkg/events"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	rabbitMQChannel := getRabbitMQChannel()

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("ProductCreated", &handler.ProductCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createProductUseCase := NewCreateProductUseCase(db, eventDispatcher)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	productHandler := ProductHandler(db, eventDispatcher)
	webserver.AddHandler("POST", "/", echo.HandlerFunc(productHandler.Create))
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

}
func getRabbitMQChannel() *amqp.Channel {
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

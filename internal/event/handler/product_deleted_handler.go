package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
	"github.com/vitorconti/go-products/pkg/events"
)

type ProductDeletedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewProductDeletedHandler(rabbitMQChannel *amqp.Channel) *ProductDeletedHandler {
	return &ProductDeletedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *ProductDeletedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Product deleted: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct",
		"",
		false,
		false,
		msgRabbitmq,
	)
}

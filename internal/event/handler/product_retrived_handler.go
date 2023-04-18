package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
	"github.com/vitorconti/go-products/pkg/events"
)

type ProductRetrivedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewProductRetrivedHandler(rabbitMQChannel *amqp.Channel) *ProductRetrivedHandler {
	return &ProductRetrivedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *ProductRetrivedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Product created: %v", event.GetPayload())
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

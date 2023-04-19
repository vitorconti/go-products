package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/streadway/amqp"
	"github.com/vitorconti/go-products/pkg/events"
)

type ProductUpdatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewProductUpdatedHandler(rabbitMQChannel *amqp.Channel) *ProductUpdatedHandler {
	return &ProductUpdatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *ProductUpdatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
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

package event

import "time"

type ProductDeleted struct {
	Name    string
	Payload interface{}
}

func NewProductDeleted() *ProductDeleted {
	return &ProductDeleted{
		Name: "ProductDeleted",
	}
}

func (e *ProductDeleted) GetName() string {
	return e.Name
}

func (e *ProductDeleted) GetPayload() interface{} {
	return e.Payload
}

func (e *ProductDeleted) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *ProductDeleted) GetDateTime() time.Time {
	return time.Now()
}

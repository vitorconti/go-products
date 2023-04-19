package event

import "time"

type ProductUpdated struct {
	Name    string
	Payload interface{}
}

func NewProductUpdated() *ProductUpdated {
	return &ProductUpdated{
		Name: "ProductUpdated",
	}
}

func (e *ProductUpdated) GetName() string {
	return e.Name
}

func (e *ProductUpdated) GetPayload() interface{} {
	return e.Payload
}

func (e *ProductUpdated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *ProductUpdated) GetDateTime() time.Time {
	return time.Now()
}

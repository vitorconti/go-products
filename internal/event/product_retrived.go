package event

import "time"

type ProductRetrived struct {
	Name    string
	Payload interface{}
}

func NewProductRetrived() *ProductRetrived {
	return &ProductRetrived{
		Name: "ProductRetrived",
	}
}

func (e *ProductRetrived) GetName() string {
	return e.Name
}

func (e *ProductRetrived) GetPayload() interface{} {
	return e.Payload
}

func (e *ProductRetrived) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *ProductRetrived) GetDateTime() time.Time {
	return time.Now()
}

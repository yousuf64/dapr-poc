package events

import (
	"errors"
	"github.com/yousuf64/go-event-sourcing/command-service/internal/aggregates/productaggregate"
)

type ProductCreated struct {
	EventRoot

	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Quantity    int32   `json:"quantity"`
}

func (pc *ProductCreated) Apply(aggregate *productaggregate.ProductAggregate, track bool) error {
	if pc.Quantity < 1 {
		return errors.New("invalid quantity")
	}
	if pc.Price < 0 {
		return errors.New("invalid price")
	}
	aggregate.Price = pc.Price
	aggregate.Quantity = pc.Quantity
	if track {
		aggregate.TrackChange(pc)
	}
	return nil
}

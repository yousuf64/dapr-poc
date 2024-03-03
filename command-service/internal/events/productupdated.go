package events

import (
	"errors"
	"github.com/yousuf64/go-event-sourcing/command-service/internal/aggregates/productaggregate"
)

type ProductUpdated struct {
	EventRoot

	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
}

func (pc *ProductUpdated) Apply(aggregate *productaggregate.ProductAggregate, track bool) error {
	if aggregate.Deactivated {
		return errors.New("product deactivated")
	}
	aggregate.Price = pc.Price
	if track {
		aggregate.TrackChange(pc)
	}
	return nil
}

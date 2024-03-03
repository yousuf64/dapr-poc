package events

import (
	"errors"
	"github.com/yousuf64/go-event-sourcing/command-service/internal/aggregates/productaggregate"
)

type ProductSold struct {
	EventRoot

	Quantity int32 `json:"quantity"`
}

func (pc *ProductSold) Apply(aggregate *productaggregate.ProductAggregate, track bool) error {
	if pc.Quantity < 1 {
		return errors.New("invalid quantity")
	}
	if aggregate.Deactivated {
		return errors.New("product deactivated")
	}
	if aggregate.Quantity-pc.Quantity < 0 {
		return errors.New("insufficient stocks")
	}
	aggregate.Quantity -= pc.Quantity
	if track {
		aggregate.TrackChange(pc)
	}
	return nil
}

package events

import (
	"errors"
	"github.com/yousuf64/go-event-sourcing/command-service/internal/aggregates/productaggregate"
)

type ProductRestocked struct {
	EventRoot

	Quantity int32 `json:"quantity"`
}

func (pc *ProductRestocked) Apply(aggregate *productaggregate.ProductAggregate, track bool) error {
	if pc.Quantity < 1 {
		return errors.New("invalid quantity")
	}
	aggregate.Quantity += pc.Quantity
	if track {
		aggregate.TrackChange(pc)
	}
	return nil
}

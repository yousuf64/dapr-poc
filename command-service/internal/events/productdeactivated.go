package events

import (
	"errors"
	"github.com/yousuf64/go-event-sourcing/command-service/internal/aggregates/productaggregate"
)

type ProductDeactivated struct {
	EventRoot
}

func (pc *ProductDeactivated) Apply(aggregate *productaggregate.ProductAggregate, track bool) error {
	if aggregate.Deactivated {
		return errors.New("product deactivated already")
	}
	aggregate.Deactivated = true
	if track {
		aggregate.TrackChange(pc)
	}
	return nil
}

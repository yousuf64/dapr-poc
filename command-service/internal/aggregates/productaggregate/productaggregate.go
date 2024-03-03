package productaggregate

import (
	"context"
	"errors"
	"github.com/yousuf64/go-event-sourcing/command-service/internal/aggregates"
	"github.com/yousuf64/go-event-sourcing/command-service/pkg/store"
)

type ProductAggregate struct {
	Price       float32
	Quantity    int32
	Deactivated bool

	store   *store.Store
	changes []aggregates.Event[ProductAggregate]
}

func New(s *store.Store) *ProductAggregate {
	return &ProductAggregate{
		Price:       0,
		Quantity:    0,
		Deactivated: false,
		store:       s,
		changes:     []aggregates.Event[ProductAggregate]{},
	}
}

func (pa *ProductAggregate) TrackChange(evt aggregates.Event[ProductAggregate]) {
	pa.changes = append(pa.changes, evt)
}

func (pa *ProductAggregate) Save(ctx context.Context) error {
	if len(pa.changes) == 0 {
		return errors.New("no changes to save")
	}

	for _, event := range pa.changes {
		err := pa.store.Save(ctx, event.GetId(), event.GetPartitionKey(), event)
		if err != nil {
			return err
		}
	}

	pa.changes = []aggregates.Event[ProductAggregate]{}
	return nil
}

func FromStore(s *store.Store) (*ProductAggregate, error) {
	return &ProductAggregate{}, nil
}

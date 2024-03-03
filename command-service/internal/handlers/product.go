package handlers

import (
	"context"
	"github.com/yousuf64/go-event-sourcing/command-service/internal"
	"github.com/yousuf64/go-event-sourcing/command-service/internal/aggregates/productaggregate"
	"github.com/yousuf64/go-event-sourcing/command-service/internal/commands"
	"github.com/yousuf64/go-event-sourcing/command-service/internal/events"
	"github.com/yousuf64/go-event-sourcing/command-service/pkg/idprovider"
	"github.com/yousuf64/go-event-sourcing/command-service/pkg/store"
	"time"
)

type ProductHandlers struct {
	idprovider *idprovider.IdProvider
	store      *store.Store
}

func NewProductHandlers(idp *idprovider.IdProvider, s *store.Store) *ProductHandlers {
	return &ProductHandlers{
		idprovider: idp,
		store:      s,
	}
}

type EventType string

const ProductCreated EventType = "ProductCreated"

func (p *ProductHandlers) CreateProduct(ctx context.Context, cmd commands.CreateProduct) (out commands.CreateProductOut, err error) {
	id := p.idprovider.Generate()

	event := events.ProductCreated{
		EventRoot: events.EventRoot{
			Id:          id,
			Kind:        internal.ProductCreated,
			Version:     1,
			Aggregate:   internal.ProductAggregate,
			AggregateId: p.idprovider.Generate(),
			BucketId:    id,
			CreatedAt:   time.Now(),
		},
		Name:        cmd.Name,
		Description: cmd.Description,
		Price:       cmd.Price,
		Quantity:    cmd.Quantity,
	}

	productAggregate := productaggregate.New(p.store)
	err = event.Apply(productAggregate, true)
	if err != nil {
		return out, err
	}

	err = productAggregate.Save(ctx)
	if err != nil {
		return out, err
	}

	return commands.CreateProductOut{
		ProductId: event.AggregateId,
	}, nil
}

func (p *ProductHandlers) UpdateProduct(ctx context.Context, cmd commands.UpdateProduct) error {
	event := events.ProductUpdated{
		EventRoot: events.EventRoot{
			Id:          p.idprovider.Generate(),
			Kind:        internal.ProductCreated,
			Version:     1,
			Aggregate:   internal.ProductAggregate,
			AggregateId: cmd.ProductId,
			BucketId:    cmd.BucketId,
			CreatedAt:   time.Now(),
		},
		Name:        cmd.Name,
		Description: cmd.Description,
		Price:       cmd.Price,
	}

	productAggregate, err := productaggregate.FromStore(p.store)
	if err != nil {
		return err
	}
	err = event.Apply(productAggregate, true)
	if err != nil {
		return err
	}
	err = productAggregate.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductHandlers) DeleteProduct(ctx context.Context, cmd commands.DeleteProduct) error {
	event := events.ProductDeactivated{
		EventRoot: events.EventRoot{
			Id:          p.idprovider.Generate(),
			Kind:        internal.ProductCreated,
			Version:     1,
			Aggregate:   internal.ProductAggregate,
			AggregateId: cmd.ProductId,
			BucketId:    cmd.BucketId,
			CreatedAt:   time.Now(),
		},
	}

	productAggregate, err := productaggregate.FromStore(p.store)
	if err != nil {
		return err
	}
	err = event.Apply(productAggregate, true)
	if err != nil {
		return err
	}
	err = productAggregate.Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

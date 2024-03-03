package events

import (
	"github.com/yousuf64/go-event-sourcing/command-service/internal"
	"time"
)

type EventRoot struct {
	Id          string             `json:"id"`
	Kind        internal.Kind      `json:"kind"`
	Version     int                `json:"version"`
	Aggregate   internal.Aggregate `json:"aggregate"`
	AggregateId string             `json:"aggregateId"`
	BucketId    string             `json:"bucketId"`
	CreatedAt   time.Time          `json:"createdAt"`
}

func (e *EventRoot) GetId() string {
	return e.Id
}

func (e *EventRoot) GetPartitionKey() string {
	return e.BucketId
}

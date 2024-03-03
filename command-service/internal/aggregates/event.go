package aggregates

type Event[TAggregate any] interface {
	Apply(aggregate *TAggregate, track bool) error
	GetId() string
	GetPartitionKey() string
	//Kind() string
	//Version() int
	//Aggregate() string
	//AggregateId() string
	//CreatedAt() time.Time
}

//func (e *EventRoot[TAggregate]) Apply(aggregate *TAggregate, track bool) error {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *EventRoot[TAggregate]) GetId() string {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *EventRoot[TAggregate]) GetKind() string {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *EventRoot[TAggregate]) GetVersion() int {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *EventRoot[TAggregate]) GetAggregate() string {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *EventRoot[TAggregate]) GetAggregateId() string {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (e *EventRoot[TAggregate]) GetCreatedAt() time.Time {
//	//TODO implement me
//	panic("implement me")
//}

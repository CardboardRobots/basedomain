package domain

import (
	"context"
	"time"
)

type MongoEvent interface {
	Type() string
}

type EventService[T AggregateRoot[U], U ~string] interface {
	Append(ctx context.Context, options AppendOptions, events ...AppendEvent[U]) error
	GetByID(ctx context.Context, id U) (t T, version Version, err error)
	GetByIDAndName(ctx context.Context, id U, name AggregateName) (t T, version Version, err error)
}

type AggregateRoot[U ~string] interface {
	ID() U
	Apply(any) // The method to project an Event onto the Aggregate
}

type AppendOptions struct {
	RequestID     RequestID     // The ID of the Request, used for idempotency
	CorrelationID CorrelationID // The ID of the entire Command and Event chain
	UserID        UserID        // The ID of the User issuing the Command
	Time          time.Time     // The Time this Event was created
}

type AppendEvent[U ~string] struct {
	EventID     EventID     // The ID of the Event, use for de-duplication
	AggregateID U           // The ID of the Aggregate to which this Event applies
	Version     Version     // The Version of the Aggregate
	PartitionID PartitionID // The ID to partition Events
	Data        BaseEvent   // The Data for the Event
}

type BaseEvent interface {
	AggregateName() string
	EventName() string
	Schema() string
}

func NewAppendEvent[U ~string](
	aggregateID U,
	version Version,
	partitionID PartitionID,
	data BaseEvent,
) AppendEvent[U] {
	return AppendEvent[U]{
		AggregateID: aggregateID,
		Version:     version + 1,
		PartitionID: partitionID,
		Data:        data,
	}
}

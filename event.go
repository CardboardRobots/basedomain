package domain

import "time"

type Event struct {
	GlobalVersion GlobalVersion // The global version of the Event
	ID            EventID       // The ID of the Event, used for de-duplication
	AggregateID   AggregateID   // The ID of the Aggregate to which this Event applies
	Version       Version       // The Version of the Aggregate
	PartitionID   PartitionID   // The ID to partition Events
	AggregateName AggregateName // The Name of the Aggregate
	Name          EventName     // The Name of the Event
	Schema        Schema        // The Schema for the internal data
	Metadata      EventMetadata // The Metadata for the Event
	Data          []byte        // The internal data of the Event
}

type EventMetadata struct {
	CorrelationID CorrelationID // The ID of the entire Command and Event chain
	UserID        UserID        // The ID of the User issuing the Command
	Time          time.Time     // The Time this Event was created
}

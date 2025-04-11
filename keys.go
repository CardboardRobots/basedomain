package domain

import "github.com/google/uuid"

type (
	AggregateID   string
	CorrelationID string
	UserID        string
	EventID       string
	PartitionID   int32
	Schema        string
	GlobalVersion int
	RequestID     string
)

func NewAggregateID() AggregateID     { return AggregateID(uuid.NewString()) }
func NewCorrelationID() CorrelationID { return CorrelationID(uuid.NewString()) }
func NewUserID() UserID               { return UserID(uuid.NewString()) }
func NewEventID() EventID             { return EventID(uuid.NewString()) }
func NewSchema() Schema               { return Schema(uuid.NewString()) }
func NewRequestID() RequestID         { return RequestID(uuid.NewString()) }

func (i AggregateID) String() string   { return string(i) }
func (i CorrelationID) String() string { return string(i) }
func (i UserID) String() string        { return string(i) }
func (i EventID) String() string       { return string(i) }
func (i Schema) String() string        { return string(i) }
func (i RequestID) String() string     { return string(i) }

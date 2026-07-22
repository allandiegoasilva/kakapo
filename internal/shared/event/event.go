package event

import (
	"time"

	"github.com/google/uuid"
)

type DomainEvent[T any] struct {
	ID          string
	AggregateID string
	EventType   string
	OccurredAt  time.Time
	Payload     T
}

type NewDomainInput[T any] struct {
	EventName   string
	AggregateID string
	Payload     T
}

func NewDomainEvent[T any](input NewDomainInput[T]) *DomainEvent[T] {
	return &DomainEvent[T]{
		ID:          uuid.NewString(),
		AggregateID: input.AggregateID,
		EventType:   input.EventName,
		OccurredAt:  time.Now(),
		Payload:     input.Payload,
	}
}

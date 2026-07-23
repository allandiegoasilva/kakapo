package project

import (
	"kakapo/internal/shared/event"
)

const (
	CreatedEventName = "project.created"
	UpdatedEventName = "project.updated"
)

type (
	CreatedEvent = event.DomainEvent[Project]
	UpdatedEvent = event.DomainEvent[Project]
)

func NewCreatedEvent(p *Project) *CreatedEvent {
	return event.NewDomainEvent(event.NewDomainInput[Project]{
		EventName:   CreatedEventName,
		AggregateID: p.ID,
		Payload:     *p,
	})
}

func NewUpdatedEvent(p *Project) *UpdatedEvent {
	return event.NewDomainEvent(event.NewDomainInput[Project]{
		EventName:   UpdatedEventName,
		AggregateID: p.ID,
		Payload:     *p,
	})
}

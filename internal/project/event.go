package project

import "kakapo/internal/shared"

const (
	CreatedEventName = "project.created"
	UpdatedEventName = "project.updated"
)

type (
	CreatedEvent = shared.DomainEvent[Project]
	UpdatedEvent = shared.DomainEvent[Project]
)

func NewCreatedEvent(p *Project) *CreatedEvent {
	return shared.NewDomainEvent(shared.NewDomainInput[Project]{
		EventName:   CreatedEventName,
		AggregateID: p.Id,
		Payload:     *p,
	})
}

func NewUpdatedEvent(p *Project) *UpdatedEvent {
	return shared.NewDomainEvent(shared.NewDomainInput[Project]{
		EventName:   UpdatedEventName,
		AggregateID: p.Id,
		Payload:     *p,
	})
}

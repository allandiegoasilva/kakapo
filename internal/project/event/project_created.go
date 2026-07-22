package event

import (
	"kakapo/internal/project/entity"
	sharedevent "kakapo/internal/shared/event"
)

const ProjectCreatedEventName = "project.created"

type ProjectCreatedEvent = sharedevent.DomainEvent[entity.Project]

func NewProjectCreatedEvent(project *entity.Project) *ProjectCreatedEvent {
	return sharedevent.NewDomainEvent(sharedevent.NewDomainInput[entity.Project]{
		EventName:   ProjectCreatedEventName,
		AggregateID: project.ID,
		Payload:     *project,
	})
}

package project

import (
	shared "kakapo/internal/shared/event"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	events []shared.DomainEvent[any]
}

type ProjectCreateInput struct {
	Name        string
	Description string
}

func NewProject(in ProjectCreateInput) *Project {
	now := time.Now()
	return &Project{
		ID:          uuid.NewString(),
		Name:        in.Name,
		Description: in.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
		events:      []shared.DomainEvent[any]{},
	}
}

func (p *Project) Raise(e shared.DomainEvent[any]) {
	p.events = append(p.events, e)
}

func (p *Project) PullEvents() []shared.DomainEvent[any] {
	events := p.events
	p.events = nil
	return events
}

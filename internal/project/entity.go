package project

import (
	"kakapo/internal/shared"
	"kakapo/internal/shared/vo"
	"time"

	"github.com/google/uuid"
)

type Project struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Slug        string    `json:"slug"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	events []shared.DomainEvent[any]
}

type ProjectCreateInput struct {
	Name        string
	Description string
}

func NewProject(i ProjectCreateInput) (*Project, error) {
	now := time.Now()

	slug, err := vo.NewSlugVO(i.Name)

	if err != nil {
		return nil, err
	}

	description := vo.NewCleanStringVO(i.Description, vo.CleanStringOptions{
		MaxLength: 350,
	}).Value()

	return &Project{
		Id:          uuid.NewString(),
		Name:        vo.NewCleanStringVO(i.Name).Value(),
		Description: description,
		Slug:        slug.Value(),
		CreatedAt:   now,
		UpdatedAt:   now,
		events:      []shared.DomainEvent[any]{},
	}, nil
}

func (p *Project) Raise(e shared.DomainEvent[any]) {
	p.events = append(p.events, e)
}

func (p *Project) PullEvents() []shared.DomainEvent[any] {
	events := p.events
	p.events = nil
	return events
}

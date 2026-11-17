package project

type ProjectRepository interface {
	Create(p *Project) error
	Update(id string, p *Project) error
}

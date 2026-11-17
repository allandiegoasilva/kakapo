package project

type ProjectService struct {
	repository *ProjectRepository
}

type ProjectRepositoryInput struct {
	Repository *ProjectRepository
}

func NewService(i ProjectRepositoryInput) *ProjectService {
	return &ProjectService{
		repository: i.Repository,
	}
}

func (p *ProjectService) Create(i ProjectCreateInput) (*Project, error) {
	project, err := NewProject(i)

	if err != nil {
		return nil, err
	}

	return project, nil
}

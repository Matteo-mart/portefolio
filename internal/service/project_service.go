package service

import (
	"portefolio/internal/domain"
	"portefolio/internal/repository"
	"time"
)

type ProjectService struct {
	repo repository.ProjectRepository
}

// NewProjectService crée un nouveau service de projets
func NewProjectService(repo repository.ProjectRepository) *ProjectService {
	return &ProjectService{repo: repo}
}

// CreateProject crée un nouveau projet avec validation
func (s *ProjectService) CreateProject(req *domain.CreateProjectRequest, images []string) (int64, error) {
	// Validation
	if err := req.Validate(); err != nil {
		return 0, err
	}

	// Parser la date
	dateCreation, err := time.Parse("2006-01-02", req.DateCreation)
	if err != nil {
		return 0, domain.ErrDateRequired
	}

	// Créer le projet
	project := &domain.Project{
		Titre:        req.Titre,
		DateCreation: dateCreation,
		Description:  req.Description,
		Technologie:  req.Technologie,
		Explication:  req.Explication,
		Probleme:     req.Probleme,
		Solution:     req.Solution,
		UrlSource:    req.UrlSource,
	}

	return s.repo.Create(project, images)
}

// GetProject récupère un projet par son ID
func (s *ProjectService) GetProject(id int64) (*domain.Project, error) {
	if id <= 0 {
		return nil, domain.ErrInvalidID
	}
	return s.repo.GetByID(id)
}

// GetAllProjects récupère tous les projets
func (s *ProjectService) GetAllProjects() ([]*domain.Project, error) {
	return s.repo.GetAll()
}

// UpdateProject met à jour un projet
func (s *ProjectService) UpdateProject(id int64, req *domain.UpdateProjectRequest) error {
	if id <= 0 {
		return domain.ErrInvalidID
	}

	// Validation basique
	if req.Titre == "" {
		return domain.ErrTitreRequired
	}

	return s.repo.Update(id, req)
}

// DeleteProject supprime définitivement un projet
func (s *ProjectService) DeleteProject(id int64) error {
	if id <= 0 {
		return domain.ErrInvalidID
	}
	return s.repo.Delete(id)
}

// MoveToTrash déplace un projet dans la corbeille
func (s *ProjectService) MoveToTrash(id int64) error {
	if id <= 0 {
		return domain.ErrInvalidID
	}
	return s.repo.MoveToTrash(id)
}

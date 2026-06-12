package service

import (
	"portefolio/internal/domain"
	"portefolio/internal/repository"
)

type TechnologyService struct {
	repo repository.TechnologyRepository
}

// NewTechnologyService crée un nouveau service de technologies
func NewTechnologyService(repo repository.TechnologyRepository) *TechnologyService {
	return &TechnologyService{repo: repo}
}

// CreateTechnology crée une nouvelle technologie
func (s *TechnologyService) CreateTechnology(req *domain.CreateTechnologyRequest) (int64, error) {
	// Validation
	if err := req.Validate(); err != nil {
		return 0, err
	}

	tech := &domain.Technology{
		Nom:       req.Nom,
		Icone:     req.Icone,
		UrlSource: req.UrlSource,
	}

	return s.repo.Create(tech)
}

// GetTechnology récupère une technologie par son ID
func (s *TechnologyService) GetTechnology(id int64) (*domain.Technology, error) {
	if id <= 0 {
		return nil, domain.ErrInvalidID
	}
	return s.repo.GetByID(id)
}

// GetAllTechnologies récupère toutes les technologies
func (s *TechnologyService) GetAllTechnologies() ([]*domain.Technology, error) {
	return s.repo.GetAll()
}

// UpdateTechnology met à jour une technologie
func (s *TechnologyService) UpdateTechnology(id int64, req *domain.UpdateTechnologyRequest) error {
	if id <= 0 {
		return domain.ErrInvalidID
	}

	if req.Nom == "" {
		return domain.ErrNomRequired
	}

	return s.repo.Update(id, req)
}

// DeleteTechnology supprime définitivement une technologie
func (s *TechnologyService) DeleteTechnology(id int64) error {
	if id <= 0 {
		return domain.ErrInvalidID
	}
	return s.repo.Delete(id)
}

// MoveToTrash déplace une technologie dans la corbeille
func (s *TechnologyService) MoveToTrash(id int64) error {
	if id <= 0 {
		return domain.ErrInvalidID
	}
	return s.repo.MoveToTrash(id)
}

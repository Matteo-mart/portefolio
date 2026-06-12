package service

import (
	"portefolio/internal/domain"
	"portefolio/internal/repository"
)

type CorbeilleService struct {
	repo repository.CorbeilleRepository
}

// NewCorbeilleService crée un nouveau service de corbeille
func NewCorbeilleService(repo repository.CorbeilleRepository) *CorbeilleService {
	return &CorbeilleService{repo: repo}
}

// GetProjects récupère tous les projets dans la corbeille
func (s *CorbeilleService) GetProjects() ([]*domain.CorbeilleEntry, error) {
	return s.repo.GetProjects()
}

// GetTechnologies récupère toutes les technologies dans la corbeille
func (s *CorbeilleService) GetTechnologies() ([]*domain.CorbeilleTech, error) {
	return s.repo.GetTechnologies()
}

// RestoreProject restaure un projet depuis la corbeille
func (s *CorbeilleService) RestoreProject(id int64) error {
	if id <= 0 {
		return domain.ErrInvalidID
	}
	return s.repo.RestoreProject(id)
}

// RestoreTechnology restaure une technologie depuis la corbeille
func (s *CorbeilleService) RestoreTechnology(id int64) error {
	if id <= 0 {
		return domain.ErrInvalidID
	}
	return s.repo.RestoreTechnology(id)
}

// DeleteProject supprime définitivement un projet de la corbeille
func (s *CorbeilleService) DeleteProject(id int64) error {
	if id <= 0 {
		return domain.ErrInvalidID
	}
	return s.repo.DeleteProject(id)
}

// DeleteTechnology supprime définitivement une technologie de la corbeille
func (s *CorbeilleService) DeleteTechnology(id int64) error {
	if id <= 0 {
		return domain.ErrInvalidID
	}
	return s.repo.DeleteTechnology(id)
}

// EmptyTrash vide complètement la corbeille
func (s *CorbeilleService) EmptyTrash() error {
	return s.repo.EmptyTrash()
}

package service

import (
	"portefolio/internal/domain"
	"portefolio/internal/repository"
)

type ContactService struct {
	repo repository.ContactRepository
}

// NewContactService crée un nouveau service de contacts
func NewContactService(repo repository.ContactRepository) *ContactService {
	return &ContactService{repo: repo}
}

// GetContact récupère un contact par son ID
func (s *ContactService) GetContact(id int64) (*domain.Contact, error) {
	if id <= 0 {
		return nil, domain.ErrInvalidID
	}
	return s.repo.GetByID(id)
}

// GetFirstContact récupère le premier contact
func (s *ContactService) GetFirstContact() (*domain.Contact, error) {
	return s.repo.GetFirst()
}

// UpdateContact met à jour un contact
func (s *ContactService) UpdateContact(id int64, req *domain.UpdateContactRequest) error {
	if id <= 0 {
		return domain.ErrInvalidID
	}

	// Validation
	if err := req.Validate(); err != nil {
		return err
	}

	return s.repo.Update(id, req)
}

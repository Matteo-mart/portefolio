package repository

import (
	"database/sql"
	"portefolio/internal/domain"
)

// ProjectRepository définit les opérations sur les projets
type ProjectRepository interface {
	Create(project *domain.Project, images []string) (int64, error)
	GetByID(id int64) (*domain.Project, error)
	GetAll() ([]*domain.Project, error)
	Update(id int64, req *domain.UpdateProjectRequest) error
	Delete(id int64) error
	MoveToTrash(id int64) error
	GetImages(projectID int64) ([]string, error)
	AddImage(projectID int64, imageURL string) error
}

// TechnologyRepository définit les opérations sur les technologies
type TechnologyRepository interface {
	Create(tech *domain.Technology) (int64, error)
	GetByID(id int64) (*domain.Technology, error)
	GetAll() ([]*domain.Technology, error)
	Update(id int64, req *domain.UpdateTechnologyRequest) error
	Delete(id int64) error
	MoveToTrash(id int64) error
}

// ContactRepository définit les opérations sur les contacts
type ContactRepository interface {
	GetByID(id int64) (*domain.Contact, error)
	GetFirst() (*domain.Contact, error)
	Update(id int64, req *domain.UpdateContactRequest) error
}

// CorbeilleRepository définit les opérations sur la corbeille
type CorbeilleRepository interface {
	GetProjects() ([]*domain.CorbeilleEntry, error)
	GetTechnologies() ([]*domain.CorbeilleTech, error)
	RestoreProject(id int64) error
	RestoreTechnology(id int64) error
	DeleteProject(id int64) error
	DeleteTechnology(id int64) error
	EmptyTrash() error
}

// Repositories regroupe tous les repositories
type Repositories struct {
	Project     ProjectRepository
	Technology  TechnologyRepository
	Contact     ContactRepository
	Corbeille   CorbeilleRepository
}

// NewRepositories crée une nouvelle instance de Repositories
func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Project:     NewProjectRepository(db),
		Technology:  NewTechnologyRepository(db),
		Contact:     NewContactRepository(db),
		Corbeille:   NewCorbeilleRepository(db),
	}
}

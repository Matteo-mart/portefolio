package repository

import (
	"database/sql"
	"portefolio/internal/domain"
)

type contactRepository struct {
	db *sql.DB
}

// NewContactRepository crée un nouveau repository de contacts
func NewContactRepository(db *sql.DB) ContactRepository {
	return &contactRepository{db: db}
}

// GetByID récupère un contact par son ID
func (r *contactRepository) GetByID(id int64) (*domain.Contact, error) {
	query := `SELECT id, telephone, email, linkedin, github, updated_at FROM contacts WHERE id = ?`

	contact := &domain.Contact{}
	err := r.db.QueryRow(query, id).Scan(
		&contact.ID,
		&contact.Telephone,
		&contact.Email,
		&contact.Linkedin,
		&contact.Github,
		&contact.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrContactNotFound
		}
		return nil, err
	}

	return contact, nil
}

// GetFirst récupère le premier contact (pour la page principale)
func (r *contactRepository) GetFirst() (*domain.Contact, error) {
	query := `SELECT id, telephone, email, linkedin, github, updated_at FROM contacts LIMIT 1`

	contact := &domain.Contact{}
	err := r.db.QueryRow(query).Scan(
		&contact.ID,
		&contact.Telephone,
		&contact.Email,
		&contact.Linkedin,
		&contact.Github,
		&contact.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrContactNotFound
		}
		return nil, err
	}

	return contact, nil
}

// Update met à jour un contact
func (r *contactRepository) Update(id int64, req *domain.UpdateContactRequest) error {
	query := `UPDATE contacts SET telephone = ?, email = ?, linkedin = ?, github = ? WHERE id = ?`

	result, err := r.db.Exec(query,
		req.Telephone,
		req.Email,
		req.Linkedin,
		req.Github,
		id,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrContactNotFound
	}

	return nil
}

package repository

import (
	"database/sql"
	"portefolio/internal/domain"
)

type technologyRepository struct {
	db *sql.DB
}

// NewTechnologyRepository crée un nouveau repository de technologies
func NewTechnologyRepository(db *sql.DB) TechnologyRepository {
	return &technologyRepository{db: db}
}

// Create crée une nouvelle technologie
func (r *technologyRepository) Create(tech *domain.Technology) (int64, error) {
	query := `INSERT INTO technologies (nom, icone, url_source) VALUES (?, ?, ?)`

	result, err := r.db.Exec(query, tech.Nom, tech.Icone, tech.UrlSource)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

// GetByID récupère une technologie par son ID
func (r *technologyRepository) GetByID(id int64) (*domain.Technology, error) {
	query := `SELECT id, nom, icone, url_source FROM technologies WHERE id = ?`

	tech := &domain.Technology{}
	err := r.db.QueryRow(query, id).Scan(&tech.ID, &tech.Nom, &tech.Icone, &tech.UrlSource)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrTechnologyNotFound
		}
		return nil, err
	}

	return tech, nil
}

// GetAll récupère toutes les technologies
func (r *technologyRepository) GetAll() ([]*domain.Technology, error) {
	query := `SELECT id, nom, icone, url_source FROM technologies ORDER BY nom ASC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var technologies []*domain.Technology
	for rows.Next() {
		tech := &domain.Technology{}
		if err := rows.Scan(&tech.ID, &tech.Nom, &tech.Icone, &tech.UrlSource); err != nil {
			return nil, err
		}
		technologies = append(technologies, tech)
	}

	return technologies, rows.Err()
}

// Update met à jour une technologie
func (r *technologyRepository) Update(id int64, req *domain.UpdateTechnologyRequest) error {
	query := `UPDATE technologies SET nom = ?, icone = ?, url_source = ? WHERE id = ?`

	result, err := r.db.Exec(query, req.Nom, req.Icone, req.UrlSource, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrTechnologyNotFound
	}

	return nil
}

// Delete supprime définitivement une technologie
func (r *technologyRepository) Delete(id int64) error {
	query := `DELETE FROM technologies WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrTechnologyNotFound
	}

	return nil
}

// MoveToTrash déplace une technologie dans la corbeille
func (r *technologyRepository) MoveToTrash(id int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Récupérer la technologie
	tech, err := r.GetByID(id)
	if err != nil {
		return err
	}

	// Insérer dans la corbeille
	corbeilleQuery := `INSERT INTO corbeille_technologies (tech_id, nom, icone, url_source) VALUES (?, ?, ?, ?)`
	_, err = tx.Exec(corbeilleQuery, tech.ID, tech.Nom, tech.Icone, tech.UrlSource)
	if err != nil {
		return err
	}

	// Supprimer la technologie
	if _, err := tx.Exec(`DELETE FROM technologies WHERE id = ?`, id); err != nil {
		return err
	}

	return tx.Commit()
}

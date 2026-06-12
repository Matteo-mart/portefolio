package repository

import (
	"database/sql"
	"portefolio/internal/domain"
)

type corbeilleRepository struct {
	db *sql.DB
}

// NewCorbeilleRepository crée un nouveau repository de corbeille
func NewCorbeilleRepository(db *sql.DB) CorbeilleRepository {
	return &corbeilleRepository{db: db}
}

// GetProjects récupère tous les projets dans la corbeille
func (r *corbeilleRepository) GetProjects() ([]*domain.CorbeilleEntry, error) {
	query := `SELECT id, project_id, titre, date_suppression, date_creation, description,
		technologie, explication, probleme, solution, url_source
		FROM corbeille ORDER BY date_suppression DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []*domain.CorbeilleEntry
	for rows.Next() {
		entry := &domain.CorbeilleEntry{}
		err := rows.Scan(
			&entry.ID,
			&entry.ProjectID,
			&entry.Titre,
			&entry.DateSuppression,
			&entry.DateCreation,
			&entry.Description,
			&entry.Technologie,
			&entry.Explication,
			&entry.Probleme,
			&entry.Solution,
			&entry.UrlSource,
		)
		if err != nil {
			return nil, err
		}

		// Récupérer les images
		imageQuery := `SELECT url FROM corbeille_image WHERE project_id = ?`
		imageRows, err := r.db.Query(imageQuery, entry.ProjectID)
		if err != nil {
			return nil, err
		}

		var images []string
		for imageRows.Next() {
			var url string
			if err := imageRows.Scan(&url); err != nil {
				imageRows.Close()
				return nil, err
			}
			images = append(images, url)
		}
		imageRows.Close()
		entry.Images = images

		entries = append(entries, entry)
	}

	return entries, rows.Err()
}

// GetTechnologies récupère toutes les technologies dans la corbeille
func (r *corbeilleRepository) GetTechnologies() ([]*domain.CorbeilleTech, error) {
	query := `SELECT id, tech_id, nom, icone, url_source, date_suppression
		FROM corbeille_technologies ORDER BY date_suppression DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var technologies []*domain.CorbeilleTech
	for rows.Next() {
		tech := &domain.CorbeilleTech{}
		err := rows.Scan(
			&tech.ID,
			&tech.TechID,
			&tech.Nom,
			&tech.Icone,
			&tech.UrlSource,
			&tech.DateSuppression,
		)
		if err != nil {
			return nil, err
		}
		technologies = append(technologies, tech)
	}

	return technologies, rows.Err()
}

// RestoreProject restaure un projet depuis la corbeille
func (r *corbeilleRepository) RestoreProject(id int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Récupérer les données de la corbeille
	var entry domain.CorbeilleEntry
	query := `SELECT project_id, titre, date_creation, description, technologie,
		explication, probleme, solution, url_source FROM corbeille WHERE id = ?`

	err = tx.QueryRow(query, id).Scan(
		&entry.ProjectID,
		&entry.Titre,
		&entry.DateCreation,
		&entry.Description,
		&entry.Technologie,
		&entry.Explication,
		&entry.Probleme,
		&entry.Solution,
		&entry.UrlSource,
	)
	if err != nil {
		return err
	}

	// Réinsérer dans project
	projectQuery := `INSERT INTO project
		(id, titre, date_creation, description, technologie, explication, probleme, solution, url_source)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = tx.Exec(projectQuery,
		entry.ProjectID,
		entry.Titre,
		entry.DateCreation,
		entry.Description,
		entry.Technologie,
		entry.Explication,
		entry.Probleme,
		entry.Solution,
		entry.UrlSource,
	)
	if err != nil {
		return err
	}

	// Restaurer les images
	imageQuery := `SELECT url FROM corbeille_image WHERE project_id = ?`
	rows, err := tx.Query(imageQuery, entry.ProjectID)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err != nil {
			return err
		}
		if _, err := tx.Exec(`INSERT INTO project_image (project_id, url) VALUES (?, ?)`, entry.ProjectID, url); err != nil {
			return err
		}
	}

	// Supprimer de la corbeille
	if _, err := tx.Exec(`DELETE FROM corbeille WHERE id = ?`, id); err != nil {
		return err
	}

	return tx.Commit()
}

// RestoreTechnology restaure une technologie depuis la corbeille
func (r *corbeilleRepository) RestoreTechnology(id int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Récupérer les données
	var tech domain.CorbeilleTech
	query := `SELECT tech_id, nom, icone, url_source FROM corbeille_technologies WHERE id = ?`

	err = tx.QueryRow(query, id).Scan(&tech.TechID, &tech.Nom, &tech.Icone, &tech.UrlSource)
	if err != nil {
		return err
	}

	// Réinsérer dans technologies
	techQuery := `INSERT INTO technologies (id, nom, icone, url_source) VALUES (?, ?, ?, ?)`
	_, err = tx.Exec(techQuery, tech.TechID, tech.Nom, tech.Icone, tech.UrlSource)
	if err != nil {
		return err
	}

	// Supprimer de la corbeille
	if _, err := tx.Exec(`DELETE FROM corbeille_technologies WHERE id = ?`, id); err != nil {
		return err
	}

	return tx.Commit()
}

// DeleteProject supprime définitivement un projet de la corbeille
func (r *corbeilleRepository) DeleteProject(id int64) error {
	query := `DELETE FROM corbeille WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

// DeleteTechnology supprime définitivement une technologie de la corbeille
func (r *corbeilleRepository) DeleteTechnology(id int64) error {
	query := `DELETE FROM corbeille_technologies WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

// EmptyTrash vide complètement la corbeille
func (r *corbeilleRepository) EmptyTrash() error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.Exec(`DELETE FROM corbeille`); err != nil {
		return err
	}

	if _, err := tx.Exec(`DELETE FROM corbeille_technologies`); err != nil {
		return err
	}

	return tx.Commit()
}

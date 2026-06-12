package repository

import (
	"database/sql"
	"portefolio/internal/domain"
	"time"
)

type projectRepository struct {
	db *sql.DB
}

// NewProjectRepository crée un nouveau repository de projets
func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepository{db: db}
}

// Create crée un nouveau projet avec ses images
func (r *projectRepository) Create(project *domain.Project, images []string) (int64, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	query := `INSERT INTO project
		(titre, date_creation, description, technologie, explication, probleme, solution, url_source)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := tx.Exec(query,
		project.Titre,
		project.DateCreation,
		project.Description,
		project.Technologie,
		project.Explication,
		project.Probleme,
		project.Solution,
		project.UrlSource,
	)
	if err != nil {
		return 0, err
	}

	projectID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// Ajouter les images
	if len(images) > 0 {
		imageQuery := `INSERT INTO project_image (project_id, url) VALUES (?, ?)`
		for _, imageURL := range images {
			if _, err := tx.Exec(imageQuery, projectID, imageURL); err != nil {
				return 0, err
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return projectID, nil
}

// GetByID récupère un projet par son ID avec ses images
func (r *projectRepository) GetByID(id int64) (*domain.Project, error) {
	query := `SELECT id, titre, date_creation, description, technologie, explication, probleme, solution, url_source
		FROM project WHERE id = ?`

	project := &domain.Project{}
	err := r.db.QueryRow(query, id).Scan(
		&project.ID,
		&project.Titre,
		&project.DateCreation,
		&project.Description,
		&project.Technologie,
		&project.Explication,
		&project.Probleme,
		&project.Solution,
		&project.UrlSource,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, domain.ErrProjectNotFound
		}
		return nil, err
	}

	// Récupérer les images
	images, err := r.GetImages(id)
	if err != nil {
		return nil, err
	}
	project.Images = images

	return project, nil
}

// GetAll récupère tous les projets
func (r *projectRepository) GetAll() ([]*domain.Project, error) {
	query := `SELECT id, titre, date_creation, description, technologie, explication, probleme, solution, url_source
		FROM project ORDER BY date_creation DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []*domain.Project
	for rows.Next() {
		project := &domain.Project{}
		err := rows.Scan(
			&project.ID,
			&project.Titre,
			&project.DateCreation,
			&project.Description,
			&project.Technologie,
			&project.Explication,
			&project.Probleme,
			&project.Solution,
			&project.UrlSource,
		)
		if err != nil {
			return nil, err
		}

		// Récupérer les images pour chaque projet
		images, err := r.GetImages(project.ID)
		if err != nil {
			return nil, err
		}
		project.Images = images

		projects = append(projects, project)
	}

	return projects, rows.Err()
}

// Update met à jour un projet
func (r *projectRepository) Update(id int64, req *domain.UpdateProjectRequest) error {
	query := `UPDATE project SET
		titre = ?, description = ?, technologie = ?, explication = ?,
		probleme = ?, solution = ?, url_source = ?
		WHERE id = ?`

	result, err := r.db.Exec(query,
		req.Titre,
		req.Description,
		req.Technologie,
		req.Explication,
		req.Probleme,
		req.Solution,
		req.UrlSource,
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
		return domain.ErrProjectNotFound
	}

	return nil
}

// Delete supprime définitivement un projet
func (r *projectRepository) Delete(id int64) error {
	query := `DELETE FROM project WHERE id = ?`

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return domain.ErrProjectNotFound
	}

	return nil
}

// MoveToTrash déplace un projet dans la corbeille
func (r *projectRepository) MoveToTrash(id int64) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Récupérer le projet
	project, err := r.GetByID(id)
	if err != nil {
		return err
	}

	// Insérer dans la corbeille
	corbeilleQuery := `INSERT INTO corbeille
		(project_id, titre, date_creation, description, technologie, explication, probleme, solution, url_source)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err = tx.Exec(corbeilleQuery,
		project.ID,
		project.Titre,
		project.DateCreation,
		project.Description,
		project.Technologie,
		project.Explication,
		project.Probleme,
		project.Solution,
		project.UrlSource,
	)
	if err != nil {
		return err
	}

	// Copier les images dans la corbeille
	if len(project.Images) > 0 {
		imageQuery := `INSERT INTO corbeille_image (project_id, url) VALUES (?, ?)`
		for _, imageURL := range project.Images {
			if _, err := tx.Exec(imageQuery, project.ID, imageURL); err != nil {
				return err
			}
		}
	}

	// Supprimer le projet original
	if _, err := tx.Exec(`DELETE FROM project WHERE id = ?`, id); err != nil {
		return err
	}

	return tx.Commit()
}

// GetImages récupère les images d'un projet
func (r *projectRepository) GetImages(projectID int64) ([]string, error) {
	query := `SELECT url FROM project_image WHERE project_id = ?`

	rows, err := r.db.Query(query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var images []string
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err != nil {
			return nil, err
		}
		images = append(images, url)
	}

	return images, rows.Err()
}

// AddImage ajoute une image à un projet
func (r *projectRepository) AddImage(projectID int64, imageURL string) error {
	query := `INSERT INTO project_image (project_id, url) VALUES (?, ?)`
	_, err := r.db.Exec(query, projectID, imageURL)
	return err
}

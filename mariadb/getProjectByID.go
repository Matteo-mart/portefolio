package mariadb

import (
	"database/sql"
	"fmt"
	"portefolio/models"
)

func GetProjectByID(id int) (models.Project, error) {
	var p models.Project
	err := DB.QueryRow(`SELECT id, titre, date_creation, description, technologie, explication, probleme, solution, url_source
        FROM project WHERE id = ?`, id).Scan(
		&p.ID, &p.Titre, &p.DateCreation, &p.Description,
		&p.Technologie, &p.Explication, &p.Probleme, &p.Solution, &p.UrlSource,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return p, fmt.Errorf("aucun projet trouvé avec l'id %d", id)
		}
		return p, err
	}

	rows, err := DB.Query("SELECT url FROM project_image WHERE project_id = ?", id)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var url string
			if err := rows.Scan(&url); err == nil {
				p.Images = append(p.Images, url)
			}
		}
	}

	return p, nil
}

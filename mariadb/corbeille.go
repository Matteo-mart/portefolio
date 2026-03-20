package mariadb

import (
	"fmt"
	"portefolio/models"
)

/*
Corbeille image (sélectionne url) via 'project_id'
*/
func GetCorbeilleImages(projectID int) ([]string, error) {
	rows, err := DB.Query("SELECT url FROM corbeille_image WHERE project_id = ?", projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var urls []string
	for rows.Next() {
		var url string
		if err := rows.Scan(&url); err == nil {
			urls = append(urls, url)
		}
	}
	return urls, nil
}

/*
Sélectionne le 'project_id' via son id pour le supprimer définitivement
*/
func DeleteFromCorbeille(id string) error {
	var projectID int
	err := DB.QueryRow("SELECT project_id FROM corbeille WHERE id = ?", id).Scan(&projectID)
	if err != nil {
		return fmt.Errorf("entrée introuvable : %v", err)
	}

	_, err = DB.Exec("DELETE FROM project_image WHERE project_id = ?", projectID)
	if err != nil {
		return fmt.Errorf("erreur suppression images : %v", err)
	}

	result, err := DB.Exec("DELETE FROM corbeille WHERE id = ?", id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("aucune entrée trouvée avec l'id '%s'", id)
	}
	return nil
}

/*
Sélectionne un projet via son id et le déplace en corbeille
*/
func MoveToCorbeille(id string) error {
	var projectID int
	var titre, technologie, url_source string
	var date_creation, description, explication, probleme, solution string

	err := DB.QueryRow("SELECT id, titre, date_creation, description, technologie, explication, probleme, solution, url_source FROM project WHERE id = ?", id).Scan(
		&projectID, &titre, &date_creation, &description, &technologie, &explication, &probleme, &solution, &url_source)
	if err != nil {
		return fmt.Errorf("projet introuvable : %v", err)
	}

	_, err = DB.Exec(`INSERT INTO corbeille 
        (project_id, titre, date_creation, description, technologie, explication, probleme, solution, url_source) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		projectID, titre, date_creation, description, technologie, explication, probleme, solution, url_source)
	if err != nil {
		return err
	}

	_, err = DB.Exec(`
        INSERT INTO corbeille_image (project_id, url, mime_type)
        SELECT project_id, url, mime_type FROM project_image WHERE project_id = ?`, projectID)
	if err != nil {
		return fmt.Errorf("erreur copie images : %v", err)
	}

	_, err = DB.Exec("DELETE FROM project WHERE id = ?", id)
	return err
}

/*
Sélectionne un projet à restaurer via son id depuis la corbeille
*/
func RestoreFromCorbeille(id string) error {
	var projectID int
	var titre, technologie, url_source string
	var date_creation, description, explication, probleme, solution string

	err := DB.QueryRow("SELECT project_id, titre, date_creation, description, technologie, explication, probleme, solution, url_source FROM corbeille WHERE id = ?", id).Scan(
		&projectID, &titre, &date_creation, &description, &technologie, &explication, &probleme, &solution, &url_source)
	if err != nil {
		return fmt.Errorf("entrée introuvable en corbeille : %v", err)
	}

	_, err = DB.Exec(`INSERT INTO project 
        (id, titre, date_creation, description, technologie, explication, probleme, solution, url_source) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		projectID, titre, date_creation, description, technologie, explication, probleme, solution, url_source)
	if err != nil {
		return fmt.Errorf("erreur lors de la restauration : %v", err)
	}

	_, err = DB.Exec(`
        INSERT INTO project_image (project_id, url, mime_type)
        SELECT project_id, url, mime_type FROM corbeille_image WHERE project_id = ?`, projectID)
	if err != nil {
		return fmt.Errorf("erreur restauration images : %v", err)
	}

	DB.Exec("DELETE FROM corbeille_image WHERE project_id = ?", projectID)
	_, err = DB.Exec("DELETE FROM corbeille WHERE id = ?", id)
	return err
}

/*
Supprimer tous les projets qui se trouvent dans la corbeille
*/
func ViderCorbeille() error {
	rows, err := DB.Query("SELECT project_id FROM corbeille")
	if err != nil {
		return err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var pid int
		if err := rows.Scan(&pid); err != nil {
			return err
		}
		ids = append(ids, pid)
	}

	for _, pid := range ids {
		DB.Exec("DELETE FROM project_image WHERE project_id = ?", pid)
	}

	_, err = DB.Exec("DELETE FROM corbeille")
	return err
}

/*
Permet d'afficher tous les projets qui sont dans la corbeille
et les ranger par date de suppression
*/
func GetCorbeille() ([]models.CorbeilleEntry, error) {
	rows, err := DB.Query("SELECT id, project_id, titre, date_suppression FROM corbeille ORDER BY date_suppression DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []models.CorbeilleEntry
	for rows.Next() {
		var e models.CorbeilleEntry
		if err := rows.Scan(&e.ID, &e.ProjectID, &e.Titre, &e.DateSuppression); err != nil {
			return nil, err
		}

		imgRows, err := DB.Query("SELECT url FROM corbeille_image WHERE project_id = ?", e.ProjectID)
		if err == nil {
			defer imgRows.Close()
			for imgRows.Next() {
				var url string
				if err := imgRows.Scan(&url); err == nil {
					e.Images = append(e.Images, url)
				}
			}
		}

		entries = append(entries, e)
	}
	return entries, nil
}

/*
Permet de mettre dans la corbeille les technologies
*/
func MoveToCorbeilleTech(id int) error {
	var nom, icone, url_source string
	err := DB.QueryRow("SELECT nom, icone, url_source FROM technologies WHERE id = ?", id).Scan(&nom, &icone, &url_source)
	if err != nil {
		return fmt.Errorf("technologie introuvable : %v", err)
	}

	_, err = DB.Exec(`INSERT INTO corbeille_technologies (tech_id, nom, icone, url_source) VALUES (?, ?, ?, ?)`,
		id, nom, icone, url_source)
	if err != nil {
		return fmt.Errorf("erreur insertion corbeille : %v", err)
	}

	_, err = DB.Exec("DELETE FROM technologies WHERE id = ?", id)
	return err
}

/*
Permet de restaurer les technologies depuis la corbeille
*/
func RestoreFromCorbeilleTech(id int) error {
	var techID int
	var nom, icone, url_source string

	err := DB.QueryRow("SELECT tech_id, nom, icone, url_source FROM corbeille_technologies WHERE id = ?", id).Scan(
		&techID, &nom, &icone, &url_source)
	if err != nil {
		return fmt.Errorf("entrée introuvable en corbeille : %v", err)
	}

	_, err = DB.Exec(`INSERT INTO technologies (id, nom, icone, url_source) VALUES (?, ?, ?, ?)`,
		techID, nom, icone, url_source)
	if err != nil {
		return fmt.Errorf("erreur restauration : %v", err)
	}

	_, err = DB.Exec("DELETE FROM corbeille_technologies WHERE id = ?", id)
	return err
}

/*
Permet de supprimer les technologies depuis la corbeille
*/
func DeleteFromCorbeilleTech(id int) error {
	result, err := DB.Exec("DELETE FROM corbeille_technologies WHERE id = ?", id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("aucune entrée trouvée avec l'id '%d'", id)
	}
	return nil
}

/*
Permet d'afficher les technologies dans la corbeille
*/
func GetCorbeilleTech() ([]models.CorbeilleTech, error) {
	rows, err := DB.Query("SELECT id, tech_id, nom, icone, url_source, date_suppression FROM corbeille_technologies ORDER BY date_suppression DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []models.CorbeilleTech
	for rows.Next() {
		var e models.CorbeilleTech
		if err := rows.Scan(&e.ID, &e.TechID, &e.Nom, &e.Icone, &e.UrlSource, &e.DateSuppression); err != nil {
			return nil, err
		}
		entries = append(entries, e)
	}
	return entries, nil
}

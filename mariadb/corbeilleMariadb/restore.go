package corbeillemariadb

import "fmt"

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

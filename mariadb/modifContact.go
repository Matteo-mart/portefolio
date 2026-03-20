package mariadb

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

/*
Modifier contact via id
*/
func UpdateContact(id string, telephone string, email string, linkedin string, github string) error {

	query := `
		UPDATE contacts 
		SET telephone = ?, email = ?, linkedin = ?, github = ?, updated_at = NOW() 
		WHERE id = ?`

	result, err := DB.Exec(query, telephone, email, linkedin, github, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("aucun contact trouvé avec l'id '%s' (ou aucune modification nécessaire)", id)
	}

	return nil
}

/*
Modifier projet via id
*/
func UpdateProjet(id string, Titre string, Description string, Technologie string, Explication string, Probleme string, Solution string, UrlSource string) error {

	query := `
		UPDATE project
		SET titre = ?, description = ?, technologie = ?, explication = ?, probleme = ?, solution = ?, url_Source = ?
		WHERE id = ?`

	result, err := DB.Exec(query, Titre, Description, Technologie, Explication, Probleme, Solution, UrlSource, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("aucun projet trouvé avec l'id '%s'", id)
	}

	return nil
}

/*
Ajouter une image à un projet
*/
func AddImageToProject(ID int, imagePath string) error {
	query := "INSERT INTO projet_image (project_id, url) VALUES (?, ?)"
	_, err := DB.Exec(query, ID, imagePath)
	return err
}

/*
L'insertion pour projet
*/
func InsertProject(titre, date, desc, tech, expl, prob, sol, url string) error {

	query := `INSERT INTO project (titre, date_creation, description, technologie, explication, probleme, solution, url_source) 
              VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := DB.Exec(query, titre, date, desc, tech, expl, prob, sol, url)
	return err
}

/*
Supprimer un projet
*/
func DeleteProject(id string) error {
	query := "DELETE FROM project WHERE id = ?"

	result, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("aucun projet trouvé avec le id '%s'", id)
	}

	return nil
}

package corbeillemariadb

import (
	"fmt"
	"portefolio/mariadb"
)

// var DB *sql.DB

/*
Sélectionne le 'project_id' via son id pour le supprimer définitivement
*/
func DeleteFromCorbeille(id string) error {
	var projectID int
	err := mariadb.DB.QueryRow("SELECT project_id FROM corbeille WHERE id = ?", id).Scan(&projectID)
	if err != nil {
		return fmt.Errorf("entrée introuvable : %v", err)
	}

	_, err = mariadb.DB.Exec("DELETE FROM project_image WHERE project_id = ?", projectID)
	if err != nil {
		return fmt.Errorf("erreur suppression images : %v", err)
	}

	result, err := mariadb.DB.Exec("DELETE FROM corbeille WHERE id = ?", id)
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
Supprimer tous les projets qui se trouvent dans la corbeille
*/
func ViderCorbeille() error {
	rows, err := mariadb.DB.Query("SELECT project_id FROM corbeille")
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
		mariadb.DB.Exec("DELETE FROM project_image WHERE project_id = ?", pid)
	}

	_, err = mariadb.DB.Exec("DELETE FROM corbeille")
	return err
}

/*
Permet de supprimer les technologies depuis la corbeille
*/
func DeleteFromCorbeilleTech(id int) error {
	result, err := mariadb.DB.Exec("DELETE FROM corbeille_technologies WHERE id = ?", id)
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
Supprimer un projet
*/
func DeleteProject(id string) error {
	query := "DELETE FROM project WHERE id = ?"

	result, err := mariadb.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("aucun projet trouvé avec le id '%s'", id)
	}

	return nil
}

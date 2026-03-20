package mariadb

import (
	"fmt"
)

func UpdateTechnologies(id int, nom string, icone string, url_source string) error {
	query := `
        UPDATE technologies
        SET nom = ?, icone = ?, url_source = ?
        WHERE id = ?`
	result, err := DB.Exec(query, nom, icone, url_source, id)
	if err != nil {
		return err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("aucune technologie trouvée avec l'id '%d'", id)
	}
	return nil
}

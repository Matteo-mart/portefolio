package mariadb

import "log"

func GetAllTechnologie() ([]map[string]interface{}, error) {

	query := "SELECT id, nom, icone, url_source FROM technologies"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var technologies []map[string]interface{}
	for rows.Next() {
		var id int
		var nom string
		var icone string
		var url_source string
		if err := rows.Scan(&id, &nom, &icone, &url_source); err != nil {
			log.Println("Erreur scan technologie:", err)
			continue
		}
		technologies = append(technologies, map[string]interface{}{
			"id":         id,
			"nom":        nom,
			"icone":      icone,
			"url_source": url_source,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return technologies, nil

}

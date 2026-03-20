package mariadb

import "log"

/*
Récupère tous les projets (titres et id) et
et classe par date de creation
*/
func GetAllProjects() ([]map[string]interface{}, error) {
	query := "SELECT id, titre FROM project ORDER BY date_creation DESC"
	rows, err := DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []map[string]interface{}
	for rows.Next() {
		var id int
		var titre string
		if err := rows.Scan(&id, &titre); err != nil {
			return nil, err
		}

		projects = append(projects, map[string]interface{}{
			"id":    id,
			"titre": titre,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}

/*
Récupère toutes les technologies
*/
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

/*
Récupère toutes les infos (telephone, email, linkedin, github) de contact
*/
func GetAllContact() ([]map[string]string, error) {
	rows, err := DB.Query("SELECT telephone, email, linkedin, github FROM contacts LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []map[string]string
	for rows.Next() {
		var tel, email, link, git string
		rows.Scan(&tel, &email, &link, &git)

		m := map[string]string{
			"telephone": tel,
			"email":     email,
			"linkedin":  link,
			"github":    git,
		}
		results = append(results, m)
	}
	return results, nil
}

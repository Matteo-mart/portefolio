package mariadb

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

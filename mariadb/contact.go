package mariadb

import (
	_ "github.com/go-sql-driver/mysql"
)

/*
Récupère toutes les infos (telephone, email, linkedin, github) de contact
*/
func GetContactInfo() ([]map[string]string, error) {
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

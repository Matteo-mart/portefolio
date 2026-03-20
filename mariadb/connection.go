package mariadb

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
Func pour se connecter à la base de donnée
et sélectionner la bonne database
*/
func Connection() error {
	dsn := "matteo:matteo@tcp(127.0.0.1:3306)/portefolio"
	var err error

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	_, err = DB.Exec("USE portefolio;")
	if err != nil {
		return err
	}

	return DB.Ping()
}

func InitDB(connectionString string) {
	var err error
	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

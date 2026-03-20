package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"portefolio/mariadb"
	"portefolio/models"
	"portefolio/route"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var DB *sql.DB

func main() {

	models.ClearTerminal()
	//Connexion mariadb

	if err := mariadb.Connection(); err != nil {
		log.Fatal("Erreur de connexion : ", err)
	}
	//routes
	r := mux.NewRouter()
	route.DefRoute(r)
	// models.ListRoute(r)

	fmt.Println("Serveur Go lancé sur http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", r))
}

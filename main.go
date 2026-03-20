package main

import (
	"fmt"
	"log"
	"net/http"
	"portefolio/mariadb"
	"portefolio/route"
	"portefolio/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	utils.ClearTerminal()

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

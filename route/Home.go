package route

import (
	"html/template"
	"log"
	"net/http"
	"portefolio/mariadb"
	"portefolio/utils"
)

func HandleHome(w http.ResponseWriter, r *http.Request) {
	if utils.SetupCORS(w, r) {
		return
	}
	if r.URL.Path != "/" && r.URL.Path != "/api/projects" {
		http.NotFound(w, r)
		return
	}

	projects, err := mariadb.GetAllProjects()
	if err != nil {
		log.Printf("ERREUR SQL: %s", err)
	}

	technologies, err := mariadb.GetAllTechnologie()
	if err != nil {
		log.Printf("ERREUR SQL: %s", err)
	}

	if r.URL.Path == "/api/projects" {
		w.Header().Set("Content-Type", "application/json")
	}

	tmpl, err := template.ParseFiles("templates/home.html")
	if err != nil {
		log.Printf("Erreur chargement template: %s", err)
		http.Error(w, "Erreur de rendu", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, utils.HomeData{
		Projects:     projects,
		Technologies: technologies,
	})
}

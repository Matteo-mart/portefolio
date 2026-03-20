package route

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"portefolio/mariadb"
	"portefolio/models"
	"strconv"

	"github.com/gorilla/mux"
)

func HandleProjet(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	project, err := mariadb.GetProjectByID(id)
	if err != nil {
		log.Printf("Erreur BDD pour ID %d : %v", id, err)
		http.Error(w, "Projet introuvable", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/projet.html")
	if err != nil {
		log.Printf("Erreur template : %v", err)
		http.Error(w, "Erreur serveur interne", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, project)
	if err != nil {
		log.Printf("Erreur exécution template : %v", err)
		http.Error(w, "Erreur lors du rendu", http.StatusInternalServerError)
	}
}

func GetProjectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID := vars["id"]

	var p models.Project
	err := mariadb.DB.QueryRow("SELECT id, titre, date_creation, description, technologie, explication, probleme, solution, url_source FROM project WHERE id = ?", ID).
		Scan(&p.ID, &p.Titre, &p.DateCreation, &p.Description, &p.Technologie, &p.Explication, &p.Probleme, &p.Solution, &p.UrlSource)

	if err != nil {
		log.Printf("Erreur requête projet : %v", err)
		http.Error(w, "Projet non trouvé", http.StatusNotFound)
		return
	}

	rows, err := mariadb.DB.Query("SELECT url FROM project_image WHERE project_id = ?", ID)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var imgUrl string
			if err := rows.Scan(&imgUrl); err == nil {
				p.Images = append(p.Images, imgUrl)
			}
		}
	}

	tmpl, err := template.ParseFiles("templates/projet.html")
	// tmpl, err := template.ParseFiles("/home/matteo//projet.html")

	if err != nil {
		log.Printf("Erreur parsing template : %v", err)
		http.Error(w, "Erreur serveur interne", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, p)
	if err != nil {
		log.Printf("Erreur exécution template : %v", err)
		http.Error(w, "Erreur lors du rendu", http.StatusInternalServerError)
	}
}

func HandleProjectDetaill(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "ID de projet invalide", http.StatusBadRequest)
		return
	}

	project, err := mariadb.GetProjectByID(id)
	if err != nil {
		http.Error(w, "Projet non trouvé", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/projet.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, project)
}

func HandleUpdateProjet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	if models.SetupCORS(w, r) {
		return
	}
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var c models.ProjetUpdate
	err := json.NewDecoder(r.Body).Decode(&c)

	if err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	err = mariadb.UpdateProjet(id, c.Titre, c.Description, c.Technologie, c.Explication, c.Probleme, c.Solution, c.UrlSource)
	if err != nil {
		log.Printf("id pour modif: %s", id)
		log.Printf("Erreur SQL lors de la mise à jour de [%s] : %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func HandleDeleteProject(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodDelete {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "Erreur : Le paramètre 'id' est vide", http.StatusBadRequest)
		return
	}

	err := mariadb.DeleteProject(id)
	if err != nil {
		log.Printf("Erreur SQL lors de la suppression de [%s] : %v", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

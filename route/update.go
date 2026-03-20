package route

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"portefolio/mariadb"
	"portefolio/utils"

	"github.com/gorilla/mux"
)

func HandleUpdateContact(w http.ResponseWriter, r *http.Request) {

	if utils.SetupCORS(w, r) {
		return
	}

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPut {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var c utils.ContactUpdate
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil || c.ID == "" {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	err = mariadb.UpdateContact(c.ID, c.Telephone, c.Email, c.Linkedin, c.Github)
	if err != nil {
		log.Printf("Erreur SQL lors de la mise à jour de [%s] : %v", c.ID, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func HandleContact(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/contact.html")
	if err != nil {
		http.Error(w, "Erreur lors du chargement du template : "+err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := mariadb.GetAllContact()
	if err != nil {
		http.Error(w, "Erreur base de données", http.StatusInternalServerError)
		return
	}
	log.Printf("data contact: %+v", data)
	tmpl.ExecuteTemplate(w, "contact.html", data)

}

func HandleUpdateTechnologies(w http.ResponseWriter, r *http.Request) {
	if utils.SetupCORS(w, r) {
		return
	}
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if r.Method != http.MethodPut {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var t utils.Technologie
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil || t.ID <= 0 {
		http.Error(w, "Données invalides : ID requis", http.StatusBadRequest)
		return
	}

	err = mariadb.UpdateTechnologies(t.ID, t.Nom, t.Icone, t.Url_source)
	if err != nil {
		log.Printf("Erreur SQL lors de la mise à jour de [%d] : %v", t.ID, err)
		http.Error(w, "Erreur lors de la mise à jour", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Mise à jour réussie"})
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 << 20)
	file, handler, err := r.FormFile("myImage")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération du fichier", 400)
		return
	}
	defer file.Close()

	filePath := "./templates/uploads/" + handler.Filename
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Erreur lors de la sauvegarde", 500)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Erreur lors de l écriture", 500)
		return
	}

}

func HandleUpdateProjet(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	if utils.SetupCORS(w, r) {
		return
	}
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	var c utils.ProjetUpdate
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

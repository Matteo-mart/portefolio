package route

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"portefolio/mariadb"
	"portefolio/models"
)

func HandleUpdateContact(w http.ResponseWriter, r *http.Request) {

	if models.SetupCORS(w, r) {
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

	var c models.ContactUpdate
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

	data, err := mariadb.GetContactInfo()
	if err != nil {
		http.Error(w, "Erreur base de données", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}

func HandleUpdateTechnologies(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
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

	var t models.Technologie
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

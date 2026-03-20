package route

import (
	"encoding/json"
	"log"
	"net/http"
	"portefolio/mariadb"
	"portefolio/models"
	"strconv"
)

func HandleDeleteTechnologie(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}
	if r.Method != http.MethodDelete {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	if err := mariadb.MoveToCorbeilleTech(id); err != nil {
		log.Printf("Erreur corbeille tech [%d] : %v", id, err)
		http.Error(w, "Erreur suppression", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func HandleGetCorbeilleTech(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}
	entries, err := mariadb.GetCorbeilleTech()
	if err != nil {
		http.Error(w, "Erreur récupération corbeille", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(entries)
}

func HandleRestoreCorbeilleTech(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	if err := mariadb.RestoreFromCorbeilleTech(id); err != nil {
		http.Error(w, "Erreur restauration", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func HandleDeleteDefinitiveTech(w http.ResponseWriter, r *http.Request) {
	if models.SetupCORS(w, r) {
		return
	}
	if r.Method != http.MethodDelete {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	if err := mariadb.DeleteFromCorbeilleTech(id); err != nil {
		http.Error(w, "Erreur suppression définitive", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
